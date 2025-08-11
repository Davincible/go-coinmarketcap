// Package coinmarketcap provides a comprehensive Go SDK for the CoinMarketCap API.
//
// This package offers a clean, idiomatic Go interface to interact with all CoinMarketCap API endpoints,
// including cryptocurrency data, exchange information, global metrics, and various utility functions.
//
// # Features
//
//   - Complete API coverage for all CoinMarketCap endpoints
//   - Generic type-safe responses with APIResponse[T]
//   - Built-in rate limiting and retry mechanisms
//   - Comprehensive error handling with meaningful error messages
//   - Context support for request cancellation
//   - Flexible client configuration options
//   - Production-ready with proper HTTP client settings
//
// # Quick Start
//
// Create a new client and start making requests:
//
//	client := coinmarketcap.NewClient(
//		coinmarketcap.WithAPIKey("your-api-key"),
//	)
//
//	// Get latest cryptocurrency listings
//	listings, err := client.GetCryptocurrencyListingsLatest(context.Background(), &coinmarketcap.CryptocurrencyListingsOptions{
//		Start: coinmarketcap.Int(1),
//		Limit: coinmarketcap.Int(10),
//		Convert: []string{"USD"},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, crypto := range listings.Data {
//		fmt.Printf("%s: $%f\n", crypto.Name, *crypto.Quote["USD"].Price)
//	}
//
// # Configuration Options
//
// The client can be configured with various options:
//
//	client := coinmarketcap.NewClient(
//		coinmarketcap.WithAPIKey("your-api-key"),
//		coinmarketcap.WithSandbox(true),  // Use sandbox for testing
//		coinmarketcap.WithRateLimit(60),  // Custom rate limit
//		coinmarketcap.WithHTTPClient(&http.Client{
//			Timeout: 30 * time.Second,
//		}),
//	)
//
// # Error Handling
//
// The SDK provides detailed error information through the APIError type:
//
//	if err != nil {
//		if apiErr, ok := err.(*coinmarketcap.APIError); ok {
//			if apiErr.IsRateLimit() {
//				// Handle rate limit error
//				fmt.Printf("Rate limit exceeded: %s\n", apiErr.Error())
//			} else if apiErr.IsAuthError() {
//				// Handle authentication error
//				fmt.Printf("Authentication failed: %s\n", apiErr.Error())
//			}
//		}
//	}
//
// # API Coverage
//
// The SDK provides complete coverage of CoinMarketCap API endpoints:
//
//   - Cryptocurrency endpoints (19 endpoints)
//   - Exchange endpoints (7 endpoints)
//   - Global metrics endpoints (2 endpoints)
//   - Tools and utilities (2 endpoints)
//   - Content and community endpoints (6 endpoints)
//   - Index and fear/greed endpoints (4 endpoints)
//   - Administrative endpoints (3 endpoints)
//
// # Thread Safety
//
// The client is thread-safe and can be used concurrently from multiple goroutines.
// Rate limiting is handled internally to ensure API limits are respected.
//
// # Best Practices
//
//   - Always use context for request cancellation and timeouts
//   - Handle rate limiting errors gracefully with exponential backoff
//   - Cache frequently accessed data when appropriate
//   - Use CMC IDs instead of symbols for better reliability
//   - Implement proper error handling for production use
//
// For complete documentation and examples, visit: https://github.com/tyler/go-coinmarketcap
package coinmarketcap
