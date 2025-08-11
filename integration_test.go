//go:build integration
// +build integration

package coinmarketcap

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"
)

// Integration tests that require a real CoinMarketCap API key
// Run with: go test -tags=integration
// Requires CMC_API_KEY environment variable to be set

var (
	integrationClient *Client
	testAPIKey        string
)

func init() {
	testAPIKey = os.Getenv("CMC_API_KEY")
	if testAPIKey != "" {
		integrationClient = NewClient(
			WithAPIKey(testAPIKey),
			WithRateLimit(2.0), // 2 requests per second for faster testing
		)
	}
}

func skipIfNoAPIKey(t *testing.T) {
	if testAPIKey == "" {
		t.Skip("Skipping integration test: CMC_API_KEY environment variable not set")
	}
}

func TestIntegrationCryptocurrencyMap(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := get[[]CryptocurrencyMap](integrationClient, ctx, "/v1/cryptocurrency/map", &RequestOptions[[]CryptocurrencyMap]{
		QueryParams: NewParamBuilder().
			AddInt("start", Int(1)).
			AddInt("limit", Int(5)).
			Build(),
	})

	if err != nil {
		t.Fatalf("API request failed: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Fatalf("API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
	}

	if len(resp.Data) == 0 {
		t.Fatal("Expected at least one cryptocurrency in map")
	}

	// Verify data structure
	crypto := resp.Data[0]
	if crypto.ID <= 0 {
		t.Errorf("Expected positive ID, got %d", crypto.ID)
	}
	if crypto.Name == "" {
		t.Error("Expected non-empty name")
	}
	if crypto.Symbol == "" {
		t.Error("Expected non-empty symbol")
	}
	if crypto.Slug == "" {
		t.Error("Expected non-empty slug")
	}

	t.Logf("Successfully retrieved %d cryptocurrencies, first: %s (%s)",
		len(resp.Data), crypto.Name, crypto.Symbol)
}

func TestIntegrationCryptocurrencyListingsLatest(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := get[[]CryptocurrencyListing](integrationClient, ctx, "/v1/cryptocurrency/listings/latest", &RequestOptions[[]CryptocurrencyListing]{
		QueryParams: NewParamBuilder().
			AddInt("start", Int(1)).
			AddInt("limit", Int(5)).
			AddStringSlice("convert", []string{"USD"}).
			Build(),
	})

	if err != nil {
		t.Fatalf("API request failed: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Fatalf("API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
	}

	if len(resp.Data) == 0 {
		t.Fatal("Expected at least one cryptocurrency listing")
	}

	// Verify data structure
	listing := resp.Data[0]
	if listing.ID <= 0 {
		t.Errorf("Expected positive ID, got %d", listing.ID)
	}
	if listing.Name == "" {
		t.Error("Expected non-empty name")
	}
	if listing.Symbol == "" {
		t.Error("Expected non-empty symbol")
	}

	// Check USD quote exists
	usdQuote, exists := listing.Quote["USD"]
	if !exists {
		t.Fatal("Expected USD quote to exist")
	}
	if usdQuote.Price == nil {
		t.Error("Expected price to be set")
	} else if *usdQuote.Price <= 0 {
		t.Errorf("Expected positive price, got %f", *usdQuote.Price)
	}

	t.Logf("Successfully retrieved %d listings, first: %s (%s) - $%.2f",
		len(resp.Data), listing.Name, listing.Symbol,
		func() float64 {
			if usdQuote.Price != nil {
				return *usdQuote.Price
			}
			return 0
		}())
}

func TestIntegrationCryptocurrencyQuotesLatest(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := get[map[string]CryptocurrencyQuote](integrationClient, ctx, "/v2/cryptocurrency/quotes/latest", &RequestOptions[map[string]CryptocurrencyQuote]{
		QueryParams: NewParamBuilder().
			AddStringSlice("symbol", []string{"BTC", "ETH"}).
			AddStringSlice("convert", []string{"USD"}).
			Build(),
	})

	if err != nil {
		t.Fatalf("API request failed: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Fatalf("API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
	}

	if len(resp.Data) == 0 {
		t.Fatal("Expected at least one cryptocurrency quote")
	}

	// Check BTC exists
	btcQuote, btcExists := resp.Data["BTC"]
	if !btcExists {
		t.Fatal("Expected BTC quote to exist")
	}

	if btcQuote.Name != "Bitcoin" {
		t.Errorf("Expected Bitcoin name, got %s", btcQuote.Name)
	}

	// Check USD quote for BTC
	usdQuote, exists := btcQuote.Quote["USD"]
	if !exists {
		t.Fatal("Expected USD quote for BTC")
	}
	if usdQuote.Price == nil {
		t.Error("Expected BTC price to be set")
	} else if *usdQuote.Price <= 0 {
		t.Errorf("Expected positive BTC price, got %f", *usdQuote.Price)
	}

	t.Logf("Successfully retrieved quotes for %d cryptocurrencies, BTC: $%.2f",
		len(resp.Data),
		func() float64 {
			if usdQuote.Price != nil {
				return *usdQuote.Price
			}
			return 0
		}())
}

func TestIntegrationGlobalMetricsLatest(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := get[GlobalMetrics](integrationClient, ctx, "/v1/global-metrics/quotes/latest", &RequestOptions[GlobalMetrics]{
		QueryParams: NewParamBuilder().
			AddStringSlice("convert", []string{"USD"}).
			Build(),
	})

	if err != nil {
		t.Fatalf("API request failed: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Fatalf("API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
	}

	metrics := resp.Data

	// Verify essential metrics exist
	if metrics.BtcDominance == nil {
		t.Error("Expected BTC dominance to be set")
	} else if *metrics.BtcDominance <= 0 || *metrics.BtcDominance >= 100 {
		t.Errorf("Expected reasonable BTC dominance, got %f", *metrics.BtcDominance)
	}

	if metrics.ActiveCryptocurrencies == nil {
		t.Error("Expected active cryptocurrencies count to be set")
	} else if *metrics.ActiveCryptocurrencies <= 0 {
		t.Errorf("Expected positive active cryptocurrencies count, got %d", *metrics.ActiveCryptocurrencies)
	}

	// Check USD quote
	usdQuote, exists := metrics.Quote["USD"]
	if !exists {
		t.Fatal("Expected USD quote for global metrics")
	}
	if usdQuote.MarketCap == nil {
		t.Error("Expected global market cap to be set")
	} else if *usdQuote.MarketCap <= 0 {
		t.Errorf("Expected positive global market cap, got %f", *usdQuote.MarketCap)
	}

	t.Logf("Global metrics - BTC Dominance: %.2f%%, Active Cryptos: %d, Market Cap: $%.0f",
		func() float64 {
			if metrics.BtcDominance != nil {
				return *metrics.BtcDominance
			}
			return 0
		}(),
		func() int {
			if metrics.ActiveCryptocurrencies != nil {
				return *metrics.ActiveCryptocurrencies
			}
			return 0
		}(),
		func() float64 {
			if usdQuote.MarketCap != nil {
				return *usdQuote.MarketCap
			}
			return 0
		}())
}

func TestIntegrationExchangeListingsLatest(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := get[[]ExchangeListing](integrationClient, ctx, "/v1/exchange/listings/latest", &RequestOptions[[]ExchangeListing]{
		QueryParams: NewParamBuilder().
			AddInt("start", Int(1)).
			AddInt("limit", Int(3)).
			AddStringSlice("convert", []string{"USD"}).
			Build(),
	})

	if err != nil {
		t.Fatalf("API request failed: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Fatalf("API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
	}

	if len(resp.Data) == 0 {
		t.Fatal("Expected at least one exchange listing")
	}

	// Verify data structure
	exchange := resp.Data[0]
	if exchange.ID <= 0 {
		t.Errorf("Expected positive exchange ID, got %d", exchange.ID)
	}
	if exchange.Name == "" {
		t.Error("Expected non-empty exchange name")
	}
	if exchange.Slug == "" {
		t.Error("Expected non-empty exchange slug")
	}

	t.Logf("Successfully retrieved %d exchanges, first: %s",
		len(resp.Data), exchange.Name)
}

func TestIntegrationKeyInfo(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := get[KeyInfo](integrationClient, ctx, "/v1/key/info", &RequestOptions[KeyInfo]{})

	if err != nil {
		t.Fatalf("API request failed: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Fatalf("API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
	}

	keyInfo := resp.Data

	// Verify key info structure
	if keyInfo.Plan.Name == "" {
		t.Error("Expected plan name to be set")
	}
	if keyInfo.Plan.RateLimitMinute <= 0 {
		t.Errorf("Expected positive rate limit, got %d", keyInfo.Plan.RateLimitMinute)
	}
	if keyInfo.Plan.CreditLimitDaily < 0 {
		t.Errorf("Expected non-negative daily credit limit, got %d", keyInfo.Plan.CreditLimitDaily)
	}

	t.Logf("API Key Info - Plan: %s, Rate Limit: %d/min, Daily Credits: %d/%d used",
		keyInfo.Plan.Name,
		keyInfo.Plan.RateLimitMinute,
		keyInfo.Usage.CurrentDay.CreditsUsed,
		keyInfo.Plan.CreditLimitDaily)
}

func TestIntegrationErrorHandling(t *testing.T) {
	skipIfNoAPIKey(t)

	// Test with invalid parameters to trigger API error
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := get[[]CryptocurrencyListing](integrationClient, ctx, "/v1/cryptocurrency/listings/latest", &RequestOptions[[]CryptocurrencyListing]{
		QueryParams: NewParamBuilder().
			AddInt("start", Int(-1)). // Invalid start parameter
			AddInt("limit", Int(10)).
			Build(),
	})

	// Should get an error
	if err == nil {
		t.Fatal("Expected API error for invalid parameters, got nil")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("Expected APIError, got %T: %v", err, err)
	}

	if apiErr.ErrorCode == 0 {
		t.Error("Expected non-zero error code")
	}

	if apiErr.StatusCode < 400 {
		t.Errorf("Expected HTTP error status code >= 400, got %d", apiErr.StatusCode)
	}

	t.Logf("Successfully handled API error: %s", apiErr.Error())
}

func TestIntegrationRateLimiting(t *testing.T) {
	skipIfNoAPIKey(t)

	// Create a client with very aggressive rate limiting for testing
	testClient := NewClient(
		WithAPIKey(testAPIKey),
		WithRateLimit(0.1), // 1 request per 10 seconds
	)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	start := time.Now()

	// Make first request
	_, err := get[[]CryptocurrencyMap](testClient, ctx, "/v1/cryptocurrency/map", &RequestOptions[[]CryptocurrencyMap]{
		QueryParams: NewParamBuilder().AddInt("limit", Int(1)).Build(),
	})
	if err != nil {
		t.Fatalf("First request failed: %v", err)
	}

	firstRequestTime := time.Since(start)

	// Make second request immediately - should be rate limited
	_, err = get[[]CryptocurrencyMap](testClient, ctx, "/v1/cryptocurrency/map", &RequestOptions[[]CryptocurrencyMap]{
		QueryParams: NewParamBuilder().AddInt("limit", Int(1)).Build(),
	})
	if err != nil {
		t.Fatalf("Second request failed: %v", err)
	}

	totalTime := time.Since(start)

	// Second request should take significantly longer due to rate limiting
	expectedMinDelay := 8 * time.Second // Allow some tolerance
	if totalTime < expectedMinDelay {
		t.Errorf("Rate limiting not working properly. Expected at least %v delay, got %v",
			expectedMinDelay, totalTime)
	}

	t.Logf("Rate limiting test - First request: %v, Total time: %v",
		firstRequestTime, totalTime)
}

func TestIntegrationSandboxMode(t *testing.T) {
	if testAPIKey == "" {
		t.Skip("Skipping sandbox test: CMC_API_KEY environment variable not set")
	}

	// Test sandbox mode with the test API key
	sandboxClient := NewClient(
		WithAPIKey("b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"), // Sandbox test key
		WithSandbox(true),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := get[[]CryptocurrencyMap](sandboxClient, ctx, "/v1/cryptocurrency/map", &RequestOptions[[]CryptocurrencyMap]{
		QueryParams: NewParamBuilder().
			AddInt("start", Int(1)).
			AddInt("limit", Int(5)).
			Build(),
	})

	if err != nil {
		t.Fatalf("Sandbox API request failed: %v", err)
	}

	if resp.Status.ErrorCode != 0 {
		t.Fatalf("Sandbox API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
	}

	if len(resp.Data) == 0 {
		t.Fatal("Expected at least one cryptocurrency in sandbox response")
	}

	t.Logf("Sandbox mode working - Retrieved %d cryptocurrencies", len(resp.Data))
}
