# Go CoinMarketCap SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/Davincible/go-coinmarketcap.svg)](https://pkg.go.dev/github.com/Davincible/go-coinmarketcap)
[![Go Report Card](https://goreportcard.com/badge/github.com/Davincible/go-coinmarketcap)](https://goreportcard.com/report/github.com/Davincible/go-coinmarketcap)

**📚 [GoDoc](https://pkg.go.dev/github.com/Davincible/go-coinmarketcap) | 🧪 [Testing Guide](TESTING.md) | 📖 [API Reference](docs/reference.md) | 🚀 [Examples](examples/)**

A comprehensive, production-ready Go SDK for the CoinMarketCap API. This package provides a clean, idiomatic Go interface to interact with all CoinMarketCap API endpoints with built-in rate limiting, retry mechanisms, and comprehensive error handling.

## Features

- ✅ **Complete API Coverage**: All 53+ CoinMarketCap API endpoints
- ✅ **Type-Safe**: Full Go generics support with `APIResponse[T]`
- ✅ **Production Ready**: Built-in rate limiting, retries, and timeouts
- ✅ **Error Handling**: Detailed error types with helper methods
- ✅ **Context Support**: Full context support for cancellation and timeouts
- ✅ **Flexible Configuration**: Functional options pattern for client setup
- ✅ **Zero Dependencies**: Only uses `golang.org/x/time/rate` for rate limiting
- ✅ **Well Documented**: Comprehensive GoDoc documentation

## Installation

```bash
go get github.com/Davincible/go-coinmarketcap
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/Davincible/go-coinmarketcap"
)

func main() {
    // Create client
    client := coinmarketcap.NewClient(
        coinmarketcap.WithAPIKey("your-api-key"),
    )

    // Get latest cryptocurrency listings
    listings, err := client.GetCryptocurrencyListingsLatest(context.Background(), 
        &coinmarketcap.CryptocurrencyListingsOptions{
            Start:   coinmarketcap.Int(1),
            Limit:   coinmarketcap.Int(10),
            Convert: []string{"USD"},
        })
    if err != nil {
        log.Fatal(err)
    }

    // Display results
    for _, crypto := range listings.Data {
        if quote, exists := crypto.Quote["USD"]; exists && quote.Price != nil {
            fmt.Printf("%s: $%.2f\\n", crypto.Name, *quote.Price)
        }
    }
}
```

## Configuration Options

```go
client := coinmarketcap.NewClient(
    coinmarketcap.WithAPIKey("your-api-key"),
    coinmarketcap.WithSandbox(true),              // Use sandbox for testing
    coinmarketcap.WithRateLimit(rate.Limit(2.0)), // 2 requests per second
    coinmarketcap.WithHTTPClient(&http.Client{
        Timeout: 30 * time.Second,
    }),
    coinmarketcap.WithUserAgent("MyApp/1.0"),
)
```

## API Coverage

### Cryptocurrency Endpoints (19)
- `GetCryptocurrencyMap()` - Get cryptocurrency ID mapping
- `GetCryptocurrencyInfo()` - Get metadata for cryptocurrencies  
- `GetCryptocurrencyListingsLatest()` - Latest cryptocurrency listings
- `GetCryptocurrencyListingsHistorical()` - Historical listings
- `GetCryptocurrencyListingsNew()` - Recently added cryptocurrencies
- `GetCryptocurrencyQuotesLatest()` - Latest price quotes
- `GetCryptocurrencyQuotesHistorical()` - Historical price data
- `GetCryptocurrencyMarketPairsLatest()` - Market pairs data
- `GetCryptocurrencyOHLCVLatest()` - Latest OHLCV data
- `GetCryptocurrencyOHLCVHistorical()` - Historical OHLCV data
- `GetCryptocurrencyPricePerformanceStats()` - Price performance stats
- `GetCryptocurrencyCategories()` - All cryptocurrency categories
- `GetCryptocurrencyCategory()` - Single category details
- `GetCryptocurrencyAirdrops()` - Airdrop listings
- `GetCryptocurrencyAirdrop()` - Single airdrop details
- `GetCryptocurrencyTrendingLatest()` - Trending by search volume
- `GetCryptocurrencyTrendingMostVisited()` - Most visited
- `GetCryptocurrencyTrendingGainersLosers()` - Biggest movers

### Exchange Endpoints (7)
- `GetExchangeMap()` - Exchange ID mapping
- `GetExchangeInfo()` - Exchange metadata
- `GetExchangeListingsLatest()` - Exchange listings
- `GetExchangeQuotesLatest()` - Exchange quotes
- `GetExchangeQuotesHistorical()` - Historical exchange data
- `GetExchangeMarketPairsLatest()` - Exchange market pairs
- `GetExchangeAssets()` - Exchange wallet holdings

### Global Metrics (2)
- `GetGlobalMetricsLatest()` - Latest global metrics
- `GetGlobalMetricsHistorical()` - Historical global metrics

### Tools & Utilities (8)
- `GetFiatMap()` - Supported fiat currencies
- `GetPriceConversion()` - Currency conversion
- `GetPostmanCollection()` - API Postman collection
- `GetBlockchainStatsLatest()` - Blockchain statistics
- `GetKeyInfo()` - API key usage information
- `GetIndexCMC100Latest()` - CMC 100 Index
- `GetFearAndGreedLatest()` - Fear & Greed Index
- Plus content and community endpoints

## Error Handling

The SDK provides detailed error information through the `APIError` type:

```go
if err != nil {
    if apiErr, ok := err.(*coinmarketcap.APIError); ok {
        if apiErr.IsRateLimit() {
            fmt.Println("Rate limit exceeded:", apiErr.Error())
        } else if apiErr.IsAuthError() {
            fmt.Println("Authentication failed:", apiErr.Error())
        } else if apiErr.IsPaymentRequired() {
            fmt.Println("Payment required:", apiErr.Error())
        }
    }
}
```

## Advanced Usage

### Historical Data Analysis

```go
startTime := time.Now().AddDate(0, 0, -7).Format(time.RFC3339)
endTime := time.Now().Format(time.RFC3339)

historical, err := client.GetCryptocurrencyQuotesHistorical(ctx, 
    &coinmarketcap.CryptocurrencyQuotesHistoricalOptions{
        Symbol:    []string{"BTC", "ETH"},
        TimeStart: &startTime,
        TimeEnd:   &endTime,
        Interval:  coinmarketcap.IntervalPtr(coinmarketcap.IntervalDaily),
        Convert:   []string{"USD"},
    })
```

### Market Analysis with Filters

```go
listings, err := client.GetCryptocurrencyListingsLatest(ctx,
    &coinmarketcap.CryptocurrencyListingsOptions{
        Start:        coinmarketcap.Int(1),
        Limit:        coinmarketcap.Int(50),
        Convert:      []string{"USD", "BTC"},
        Sort:         coinmarketcap.ListingSortPtr(coinmarketcap.SortMarketCap),
        SortDir:      coinmarketcap.SortDirectionPtr(coinmarketcap.SortDesc),
        MarketCapMin: coinmarketcap.Float64(1000000), // Min $1M market cap
        PriceMin:     coinmarketcap.Float64(0.01),    // Min $0.01 price
    })
```

## Best Practices

1. **Use Context**: Always pass context for timeout and cancellation support
2. **Handle Rate Limits**: Implement proper backoff strategies for rate limit errors  
3. **Cache Data**: Cache frequently accessed data to reduce API calls
4. **Use CMC IDs**: Prefer CMC IDs over symbols for better reliability
5. **Check Pointers**: Many fields are pointers - always check for nil before using
6. **Monitor Usage**: Use `GetKeyInfo()` to monitor API credit usage

## Helper Functions

The SDK provides helper functions for working with pointer types:

```go
options := &coinmarketcap.CryptocurrencyListingsOptions{
    Start:   coinmarketcap.Int(1),                    // *int
    Limit:   coinmarketcap.Int(10),                   // *int  
    Convert: []string{"USD"},                         // []string
    Sort:    coinmarketcap.ListingSortPtr(coinmarketcap.SortMarketCap), // *ListingSort
}
```

## Testing

### Unit Tests

Run the comprehensive unit test suite:

```bash
go test -v
go test -v -cover  # with coverage
```

### Integration Tests

Test against the real CoinMarketCap API:

```bash
# Set your API key
export CMC_API_KEY=your-api-key

# Run integration tests
./scripts/run-integration-tests.sh
# or
go test -tags=integration -v
```

### Sandbox Mode

Use sandbox mode for unlimited testing without consuming credits:

```go
client := coinmarketcap.NewClient(
    coinmarketcap.WithAPIKey("b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"), // Test key
    coinmarketcap.WithSandbox(true), // Uses sandbox-api.coinmarketcap.com
)
```

See [TESTING.md](TESTING.md) for detailed testing documentation.

## Examples

See the `examples/` directory for complete working examples:

- `examples/basic/` - Basic usage examples
- `examples/advanced/` - Advanced usage with error handling and filtering

## Rate Limits

Default rate limits by plan:
- Basic/Hobbyist: 30 calls/minute  
- Startup/Standard: 60 calls/minute
- Professional: 90 calls/minute
- Enterprise: 120 calls/minute

The SDK automatically handles rate limiting and retries with exponential backoff.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📚 [API Documentation](https://coinmarketcap.com/api/documentation/)
- 🐛 [Issue Tracker](https://github.com/Davincible/go-coinmarketcap/issues)
- 💬 [Discussions](https://github.com/Davincible/go-coinmarketcap/discussions)