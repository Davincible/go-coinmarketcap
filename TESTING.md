# Testing Guide

This document describes how to run the different types of tests available for the CoinMarketCap Go SDK.

## Test Types

### 1. Unit Tests (Default)

Unit tests use mock servers and don't require an API key. They test the SDK's internal logic, error handling, and data parsing.

```bash
# Run unit tests
go test -v

# Run with coverage
go test -v -cover

# Generate detailed coverage report
go test -v -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

**Features tested:**
- Client configuration and options
- Parameter building and validation
- HTTP request handling
- Error parsing and handling
- JSON unmarshaling
- Rate limiting logic
- Helper functions

### 2. Integration Tests

Integration tests make real API calls to CoinMarketCap's servers and require a valid API key.

#### Prerequisites

1. **Get an API Key**: Sign up at [CoinMarketCap API](https://coinmarketcap.com/api/)
2. **Set Environment Variable**:
   ```bash
   export CMC_API_KEY=your-api-key-here
   ```

#### Running Integration Tests

```bash
# Using the helper script (recommended)
./scripts/run-integration-tests.sh

# Or run directly with Go
go test -tags=integration -v

# Run specific integration test
go test -tags=integration -v -run TestIntegrationCryptocurrencyMap
```

**Features tested:**
- Real API connectivity
- Authentication with live API
- Response data validation
- Rate limiting behavior
- Error handling with real API errors
- Sandbox mode functionality

#### Integration Test Coverage

| Test | Endpoint | Description |
|------|----------|-------------|
| `TestIntegrationCryptocurrencyMap` | `/v1/cryptocurrency/map` | Tests basic cryptocurrency mapping |
| `TestIntegrationCryptocurrencyListingsLatest` | `/v1/cryptocurrency/listings/latest` | Tests latest cryptocurrency listings |
| `TestIntegrationCryptocurrencyQuotesLatest` | `/v2/cryptocurrency/quotes/latest` | Tests getting quotes for specific coins |
| `TestIntegrationGlobalMetricsLatest` | `/v1/global-metrics/quotes/latest` | Tests global market metrics |
| `TestIntegrationExchangeListingsLatest` | `/v1/exchange/listings/latest` | Tests exchange listings |
| `TestIntegrationKeyInfo` | `/v1/key/info` | Tests API key information |
| `TestIntegrationErrorHandling` | Various | Tests real API error responses |
| `TestIntegrationRateLimiting` | Various | Tests rate limiting behavior |
| `TestIntegrationSandboxMode` | Various | Tests sandbox environment |

## API Usage and Costs

### Credit Consumption

Integration tests are designed to be credit-efficient:
- Most tests request only 1-5 records
- Rate limiting prevents rapid API calls
- Sandbox tests use the free test key

**Estimated credit usage per full test run: ~10-15 credits**

### Rate Limiting

Integration tests use conservative rate limiting:
- Default: 1 request per minute
- Rate limiting tests use: 1 request per 10 seconds
- Prevents hitting API limits during testing

## Sandbox Testing

For unlimited testing without consuming credits:

```bash
# Set the sandbox test key
export CMC_API_KEY=b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c

# Run tests (will automatically use sandbox data)
go test -tags=integration -v -run TestIntegrationSandboxMode
```

## Continuous Integration

### GitHub Actions Example

```yaml
name: Tests

on: [push, pull_request]

jobs:
  unit-tests:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    - name: Run unit tests
      run: go test -v -cover

  integration-tests:
    runs-on: ubuntu-latest
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    steps:
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    - name: Run integration tests
      env:
        CMC_API_KEY: ${{ secrets.CMC_API_KEY }}
      run: go test -tags=integration -v
```

## Best Practices

### For Development
1. **Run unit tests frequently** - They're fast and don't consume API credits
2. **Run integration tests before releases** - Ensure real API compatibility
3. **Use sandbox mode for experimentation** - Unlimited testing without costs
4. **Monitor API usage** - Check your usage at [CoinMarketCap Account](https://coinmarketcap.com/api/account)

### For CI/CD
1. **Always run unit tests** on every commit
2. **Run integration tests** only on main branch or releases
3. **Use separate API keys** for testing (with appropriate limits)
4. **Set timeouts** to prevent hanging tests
5. **Cache test results** when appropriate

## Troubleshooting

### Common Issues

#### "CMC_API_KEY environment variable not set"
```bash
export CMC_API_KEY=your-actual-api-key
```

#### "API key invalid" errors
- Verify your API key is correct
- Check if your plan has sufficient credits
- Ensure you're using the right environment (production vs sandbox)

#### Rate limit errors
- Tests are designed to avoid this, but if it happens:
- Wait a few minutes before retrying
- Check your current API usage
- Consider upgrading your plan

#### Timeout errors
- Check your internet connection
- CoinMarketCap API might be experiencing issues
- Try increasing test timeouts

### Getting Help

1. **Check API Status**: [CoinMarketCap Status Page](https://status.coinmarketcap.com/)
2. **Review API Docs**: [CoinMarketCap API Documentation](https://coinmarketcap.com/api/documentation/)
3. **Monitor Usage**: [Your API Account](https://coinmarketcap.com/api/account)

## Test Development

### Adding New Integration Tests

1. Follow the existing pattern in `integration_test.go`
2. Use the `skipIfNoAPIKey(t)` helper
3. Set appropriate timeouts
4. Validate response structure thoroughly
5. Log meaningful information for debugging

### Example Test Structure

```go
func TestIntegrationNewEndpoint(t *testing.T) {
    skipIfNoAPIKey(t)
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    resp, err := get[ResponseType](integrationClient, ctx, "/endpoint", &RequestOptions[ResponseType]{
        QueryParams: NewParamBuilder().
            AddInt("limit", Int(5)).
            Build(),
    })
    
    if err != nil {
        t.Fatalf("API request failed: %v", err)
    }
    
    if resp.Status.ErrorCode != 0 {
        t.Fatalf("API returned error: %d - %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
    }
    
    // Validate response data
    if len(resp.Data) == 0 {
        t.Fatal("Expected data in response")
    }
    
    t.Logf("Successfully tested new endpoint")
}
```

This comprehensive testing approach ensures both the SDK's internal correctness and its compatibility with the real CoinMarketCap API.