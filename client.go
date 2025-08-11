// Package coinmarketcap provides the core HTTP client for interacting with CoinMarketCap API.
package coinmarketcap

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

// API configuration constants
const (
	DefaultBaseURL    = "https://pro-api.coinmarketcap.com"
	SandboxBaseURL    = "https://sandbox-api.coinmarketcap.com"
	DefaultTimeout    = 30 * time.Second
	DefaultRateLimit  = 30 // requests per minute for Basic plan
	DefaultRetryDelay = 1 * time.Second
	MaxRetries        = 3
)

// ClientConfig holds configuration options for the CoinMarketCap client.
type ClientConfig struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	RateLimit  rate.Limit
	Sandbox    bool
	UserAgent  string
}

// Client represents a CoinMarketCap API client with rate limiting and retry capabilities.
type Client struct {
	apiKey      string
	baseURL     string
	httpClient  *http.Client
	rateLimiter *rate.Limiter
	userAgent   string
}

// Option represents a functional option for configuring the Client.
type Option func(*ClientConfig)

// WithAPIKey sets the API key for authentication.
func WithAPIKey(apiKey string) Option {
	return func(c *ClientConfig) {
		c.APIKey = apiKey
	}
}

// WithBaseURL sets a custom base URL for the API (useful for testing).
func WithBaseURL(baseURL string) Option {
	return func(c *ClientConfig) {
		c.BaseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *ClientConfig) {
		c.HTTPClient = client
	}
}

// WithRateLimit sets a custom rate limit (requests per second).
func WithRateLimit(limit rate.Limit) Option {
	return func(c *ClientConfig) {
		c.RateLimit = limit
	}
}

// WithSandbox enables or disables sandbox mode for testing.
func WithSandbox(sandbox bool) Option {
	return func(c *ClientConfig) {
		c.Sandbox = sandbox
		if sandbox {
			c.BaseURL = SandboxBaseURL
		}
	}
}

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(userAgent string) Option {
	return func(c *ClientConfig) {
		c.UserAgent = userAgent
	}
}

// NewClient creates a new CoinMarketCap API client with the provided options.
// If no API key is provided, requests will fail with authentication errors.
func NewClient(opts ...Option) *Client {
	config := &ClientConfig{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{Timeout: DefaultTimeout},
		RateLimit:  rate.Limit(DefaultRateLimit) / 60, // convert per-minute to per-second
		UserAgent:  "go-coinmarketcap/1.0",
	}

	for _, opt := range opts {
		opt(config)
	}

	if config.Sandbox {
		config.BaseURL = SandboxBaseURL
	}

	return &Client{
		apiKey:      config.APIKey,
		baseURL:     config.BaseURL,
		httpClient:  config.HTTPClient,
		rateLimiter: rate.NewLimiter(config.RateLimit, 1),
		userAgent:   config.UserAgent,
	}
}

// RequestOptions holds optional parameters for API requests.
type RequestOptions[T any] struct {
	QueryParams url.Values
	Headers     map[string]string
}

// doRequest performs the actual HTTP request with rate limiting, retries, and error handling.
func (c *Client) doRequest(ctx context.Context, endpoint string, opts *RequestOptions[any]) (*http.Response, error) {
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}

	reqURL := c.baseURL + endpoint
	if opts != nil && opts.QueryParams != nil && len(opts.QueryParams) > 0 {
		reqURL += "?" + opts.QueryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", c.userAgent)

	if c.apiKey != "" {
		req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)
	}

	if opts != nil && opts.Headers != nil {
		for key, value := range opts.Headers {
			req.Header.Set(key, value)
		}
	}

	var resp *http.Response
	var lastErr error

	for attempt := 0; attempt <= MaxRetries; attempt++ {
		resp, lastErr = c.httpClient.Do(req)
		if lastErr != nil {
			if attempt < MaxRetries {
				time.Sleep(DefaultRetryDelay * time.Duration(attempt+1))
				continue
			}
			return nil, fmt.Errorf("request failed after %d attempts: %w", MaxRetries+1, lastErr)
		}

		if resp.StatusCode == http.StatusTooManyRequests && attempt < MaxRetries {
			resp.Body.Close()
			retryAfter := resp.Header.Get("Retry-After")
			if retryAfter != "" {
				if seconds, err := strconv.Atoi(retryAfter); err == nil {
					time.Sleep(time.Duration(seconds) * time.Second)
				} else {
					time.Sleep(DefaultRetryDelay * time.Duration(attempt+1))
				}
			} else {
				time.Sleep(DefaultRetryDelay * time.Duration(attempt+1))
			}
			continue
		}

		break
	}

	if resp.StatusCode >= 400 {
		body, _ := getResponseBody(resp)
		resp.Body.Close()

		// Try to parse the error response to get the error code
		var errorResp struct {
			Status Status `json:"status"`
		}

		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}

		if json.Unmarshal(body, &errorResp) == nil {
			apiErr.ErrorCode = errorResp.Status.ErrorCode
			if errorResp.Status.ErrorMessage != nil {
				apiErr.Message = *errorResp.Status.ErrorMessage
			}
		}

		return nil, apiErr
	}

	return resp, nil
}

// getResponseBody reads and potentially decompresses the response body
func getResponseBody(resp *http.Response) ([]byte, error) {
	var reader io.ReadCloser = resp.Body

	// Check if response is gzip compressed
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzReader.Close()
		reader = gzReader
	}

	return io.ReadAll(reader)
}

// get performs a GET request to the specified endpoint and returns a typed response.
// It handles JSON unmarshaling, error checking, and API error responses automatically.
func get[T any](c *Client, ctx context.Context, endpoint string, opts *RequestOptions[T]) (*APIResponse[T], error) {
	var reqOpts *RequestOptions[any]
	if opts != nil {
		reqOpts = &RequestOptions[any]{
			QueryParams: opts.QueryParams,
			Headers:     opts.Headers,
		}
	}

	resp, err := c.doRequest(ctx, endpoint, reqOpts)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := getResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp APIResponse[T]
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if apiResp.Status.ErrorCode != 0 {
		errorMsg := "API error"
		if apiResp.Status.ErrorMessage != nil {
			errorMsg = *apiResp.Status.ErrorMessage
		}
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			ErrorCode:  apiResp.Status.ErrorCode,
			Message:    errorMsg,
		}
	}

	return &apiResp, nil
}

// APIError represents an error response from the CoinMarketCap API.
type APIError struct {
	StatusCode int
	ErrorCode  int
	Message    string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.ErrorCode != 0 {
		return fmt.Sprintf("API error %d (HTTP %d): %s", e.ErrorCode, e.StatusCode, e.Message)
	}
	return fmt.Sprintf("HTTP error %d: %s", e.StatusCode, e.Message)
}

// IsRateLimit returns true if the error is due to rate limiting.
func (e *APIError) IsRateLimit() bool {
	return e.StatusCode == http.StatusTooManyRequests ||
		e.ErrorCode == 1008 || // API_KEY_PLAN_MINUTE_RATE_LIMIT_REACHED
		e.ErrorCode == 1009 || // API_KEY_PLAN_DAILY_RATE_LIMIT_REACHED
		e.ErrorCode == 1010 || // API_KEY_PLAN_MONTHLY_RATE_LIMIT_REACHED
		e.ErrorCode == 1011 // IP_RATE_LIMIT_REACHED
}

// IsAuthError returns true if the error is due to authentication issues.
func (e *APIError) IsAuthError() bool {
	return e.StatusCode == http.StatusUnauthorized ||
		e.ErrorCode == 1001 || // API_KEY_INVALID
		e.ErrorCode == 1002 || // API_KEY_MISSING
		e.ErrorCode == 1005 || // API_KEY_REQUIRED
		e.ErrorCode == 1007 // API_KEY_DISABLED
}

// IsPaymentRequired returns true if the error requires payment or plan upgrade.
func (e *APIError) IsPaymentRequired() bool {
	return e.StatusCode == http.StatusPaymentRequired ||
		e.ErrorCode == 1003 || // API_KEY_PLAN_REQUIRES_PAYMENT
		e.ErrorCode == 1004 // API_KEY_PLAN_PAYMENT_EXPIRED
}

// ParamBuilder provides a fluent interface for building URL query parameters.
type ParamBuilder struct {
	values url.Values
}

// NewParamBuilder creates a new parameter builder.
func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{
		values: make(url.Values),
	}
}

// Add adds a key-value pair to the parameters if the value is not empty.
func (p *ParamBuilder) Add(key, value string) *ParamBuilder {
	if value != "" {
		p.values.Add(key, value)
	}
	return p
}

// AddInt adds an integer parameter if the value is not nil.
func (p *ParamBuilder) AddInt(key string, value *int) *ParamBuilder {
	if value != nil {
		p.values.Add(key, strconv.Itoa(*value))
	}
	return p
}

// AddFloat adds a float parameter if the value is not nil.
func (p *ParamBuilder) AddFloat(key string, value *float64) *ParamBuilder {
	if value != nil {
		p.values.Add(key, strconv.FormatFloat(*value, 'f', -1, 64))
	}
	return p
}

// AddBool adds a boolean parameter if the value is not nil.
func (p *ParamBuilder) AddBool(key string, value *bool) *ParamBuilder {
	if value != nil {
		p.values.Add(key, strconv.FormatBool(*value))
	}
	return p
}

// AddStringSlice adds a comma-separated list of strings if the slice is not empty.
func (p *ParamBuilder) AddStringSlice(key string, values []string) *ParamBuilder {
	if len(values) > 0 {
		p.values.Add(key, strings.Join(values, ","))
	}
	return p
}

// AddIntSlice adds a comma-separated list of integers if the slice is not empty.
func (p *ParamBuilder) AddIntSlice(key string, values []int) *ParamBuilder {
	if len(values) > 0 {
		stringValues := make([]string, len(values))
		for i, v := range values {
			stringValues[i] = strconv.Itoa(v)
		}
		p.values.Add(key, strings.Join(stringValues, ","))
	}
	return p
}

// AddTime adds a time parameter formatted as RFC3339 if the value is not nil.
func (p *ParamBuilder) AddTime(key string, value *time.Time) *ParamBuilder {
	if value != nil {
		p.values.Add(key, value.Format(time.RFC3339))
	}
	return p
}

// Build returns the constructed URL values.
func (p *ParamBuilder) Build() url.Values {
	return p.values
}
