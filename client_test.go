package coinmarketcap

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		options []Option
		verify  func(*Client)
	}{
		{
			name: "default client",
			options: []Option{
				WithAPIKey("test-key"),
			},
			verify: func(c *Client) {
				if c.apiKey != "test-key" {
					t.Errorf("expected API key 'test-key', got %s", c.apiKey)
				}
				if c.baseURL != DefaultBaseURL {
					t.Errorf("expected base URL %s, got %s", DefaultBaseURL, c.baseURL)
				}
				if c.userAgent != "go-coinmarketcap/1.0" {
					t.Errorf("expected user agent 'go-coinmarketcap/1.0', got %s", c.userAgent)
				}
			},
		},
		{
			name: "sandbox client",
			options: []Option{
				WithAPIKey("test-key"),
				WithSandbox(true),
			},
			verify: func(c *Client) {
				if c.baseURL != SandboxBaseURL {
					t.Errorf("expected sandbox URL %s, got %s", SandboxBaseURL, c.baseURL)
				}
			},
		},
		{
			name: "custom rate limit",
			options: []Option{
				WithAPIKey("test-key"),
				WithRateLimit(rate.Limit(10.0)),
			},
			verify: func(c *Client) {
				if c.rateLimiter.Limit() != rate.Limit(10.0) {
					t.Errorf("expected rate limit 10.0, got %v", c.rateLimiter.Limit())
				}
			},
		},
		{
			name: "custom user agent",
			options: []Option{
				WithAPIKey("test-key"),
				WithUserAgent("MyApp/2.0"),
			},
			verify: func(c *Client) {
				if c.userAgent != "MyApp/2.0" {
					t.Errorf("expected user agent 'MyApp/2.0', got %s", c.userAgent)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.options...)
			tt.verify(client)
		})
	}
}

func TestAPIError(t *testing.T) {
	tests := []struct {
		name     string
		err      *APIError
		expected string
	}{
		{
			name: "API error with code",
			err: &APIError{
				StatusCode: 400,
				ErrorCode:  1001,
				Message:    "API key invalid",
			},
			expected: "API error 1001 (HTTP 400): API key invalid",
		},
		{
			name: "HTTP error without code",
			err: &APIError{
				StatusCode: 500,
				Message:    "Internal server error",
			},
			expected: "HTTP error 500: Internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("expected error message '%s', got '%s'", tt.expected, tt.err.Error())
			}
		})
	}
}

func TestAPIErrorTypes(t *testing.T) {
	tests := []struct {
		name        string
		err         *APIError
		isRateLimit bool
		isAuth      bool
		isPayment   bool
	}{
		{
			name: "rate limit error",
			err: &APIError{
				StatusCode: 429,
				ErrorCode:  1008,
			},
			isRateLimit: true,
		},
		{
			name: "auth error",
			err: &APIError{
				StatusCode: 401,
				ErrorCode:  1001,
			},
			isAuth: true,
		},
		{
			name: "payment error",
			err: &APIError{
				StatusCode: 402,
				ErrorCode:  1003,
			},
			isPayment: true,
		},
		{
			name: "generic error",
			err: &APIError{
				StatusCode: 500,
				ErrorCode:  0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.IsRateLimit() != tt.isRateLimit {
				t.Errorf("expected IsRateLimit() to be %v", tt.isRateLimit)
			}
			if tt.err.IsAuthError() != tt.isAuth {
				t.Errorf("expected IsAuthError() to be %v", tt.isAuth)
			}
			if tt.err.IsPaymentRequired() != tt.isPayment {
				t.Errorf("expected IsPaymentRequired() to be %v", tt.isPayment)
			}
		})
	}
}

func TestParamBuilder(t *testing.T) {
	builder := NewParamBuilder()

	// Test various parameter types
	builder.Add("string", "test").
		AddInt("int", Int(123)).
		AddFloat("float", Float64(12.34)).
		AddBool("bool", Bool(true)).
		AddStringSlice("slice", []string{"a", "b", "c"}).
		AddIntSlice("intslice", []int{1, 2, 3}).
		AddTime("time", &time.Time{})

	params := builder.Build()

	if params.Get("string") != "test" {
		t.Errorf("expected string param 'test', got '%s'", params.Get("string"))
	}
	if params.Get("int") != "123" {
		t.Errorf("expected int param '123', got '%s'", params.Get("int"))
	}
	if params.Get("float") != "12.34" {
		t.Errorf("expected float param '12.34', got '%s'", params.Get("float"))
	}
	if params.Get("bool") != "true" {
		t.Errorf("expected bool param 'true', got '%s'", params.Get("bool"))
	}
	if params.Get("slice") != "a,b,c" {
		t.Errorf("expected slice param 'a,b,c', got '%s'", params.Get("slice"))
	}
	if params.Get("intslice") != "1,2,3" {
		t.Errorf("expected intslice param '1,2,3', got '%s'", params.Get("intslice"))
	}

	// Test nil values are ignored
	builder2 := NewParamBuilder()
	builder2.AddInt("nil_int", nil).
		AddFloat("nil_float", nil).
		AddBool("nil_bool", nil).
		AddTime("nil_time", nil)

	params2 := builder2.Build()
	if len(params2) != 0 {
		t.Errorf("expected no parameters for nil values, got %d", len(params2))
	}

	// Test empty slices are ignored
	builder3 := NewParamBuilder()
	builder3.AddStringSlice("empty_slice", []string{}).
		AddIntSlice("empty_int_slice", []int{})

	params3 := builder3.Build()
	if len(params3) != 0 {
		t.Errorf("expected no parameters for empty slices, got %d", len(params3))
	}
}

func TestClientDoRequest(t *testing.T) {
	// Mock server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		if r.Header.Get("X-CMC_PRO_API_KEY") != "test-key" {
			t.Errorf("expected API key header, got %s", r.Header.Get("X-CMC_PRO_API_KEY"))
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("expected Accept header, got %s", r.Header.Get("Accept"))
		}
		if r.Header.Get("User-Agent") != "go-coinmarketcap/1.0" {
			t.Errorf("expected User-Agent header, got %s", r.Header.Get("User-Agent"))
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": {"test": "value"}, "status": {"error_code": 0}}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIKey("test-key"),
		WithBaseURL(server.URL),
		WithRateLimit(rate.Limit(1000)), // High rate limit for testing
	)

	ctx := context.Background()
	resp, err := get[map[string]interface{}](client, ctx, "/test", &RequestOptions[map[string]interface{}]{})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Errorf("expected error code 0, got %d", resp.Status.ErrorCode)
	}

	if data, ok := resp.Data["test"]; !ok || data != "value" {
		t.Errorf("expected data.test to be 'value', got %v", data)
	}
}

func TestClientErrorHandling(t *testing.T) {
	// Mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
			"data": null,
			"status": {
				"error_code": 1001,
				"error_message": "API key invalid"
			}
		}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIKey("invalid-key"),
		WithBaseURL(server.URL),
		WithRateLimit(rate.Limit(1000)),
	)

	ctx := context.Background()
	_, err := get[map[string]interface{}](client, ctx, "/test", &RequestOptions[map[string]interface{}]{})

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected APIError, got %T", err)
	}

	if apiErr.ErrorCode != 1001 {
		t.Errorf("expected error code 1001, got %d", apiErr.ErrorCode)
	}

	if !apiErr.IsAuthError() {
		t.Error("expected IsAuthError() to be true")
	}
}

func TestClientTimeout(t *testing.T) {
	// Mock server that delays response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": {}, "status": {"error_code": 0}}`))
	}))
	defer server.Close()

	// Create client with very short timeout
	httpClient := &http.Client{Timeout: 10 * time.Millisecond}
	client := NewClient(
		WithAPIKey("test-key"),
		WithBaseURL(server.URL),
		WithHTTPClient(httpClient),
		WithRateLimit(rate.Limit(1000)),
	)

	ctx := context.Background()
	_, err := get[map[string]interface{}](client, ctx, "/test", &RequestOptions[map[string]interface{}]{})

	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}
