# CoinMarketCap API Reference

## Overview

The CoinMarketCap API is a suite of high-performance RESTful JSON endpoints designed for cryptocurrency market data.

**Base URL**: `https://pro-api.coinmarketcap.com`

**Sandbox URL**: `https://sandbox-api.coinmarketcap.com` (for testing with mock data)

## Authentication

All requests require an API key. You can provide it in two ways:

1. **Custom Header (Preferred)**: `X-CMC_PRO_API_KEY: your-api-key`
2. **Query Parameter**: `CMC_PRO_API_KEY=your-api-key`

## Common Headers

- `Accept: application/json` (required)
- `Accept-Encoding: deflate, gzip` (recommended)

## Response Format

All endpoints return JSON with this structure:

```json
{
  "data": { ... },
  "status": {
    "timestamp": "2018-06-06T07:52:27.273Z",
    "error_code": 0,
    "error_message": null,
    "elapsed": 10,
    "credit_count": 1
  }
}
```

## Error Codes

### HTTP Status Codes
- `400` - Bad Request
- `401` - Unauthorized 
- `402` - Payment Required
- `403` - Forbidden
- `429` - Too Many Requests
- `500` - Internal Server Error

### API Error Codes
- `1001` - API_KEY_INVALID
- `1002` - API_KEY_MISSING
- `1003` - API_KEY_PLAN_REQUIRES_PAYMENT
- `1004` - API_KEY_PLAN_PAYMENT_EXPIRED
- `1005` - API_KEY_REQUIRED
- `1006` - API_KEY_PLAN_NOT_AUTHORIZED
- `1007` - API_KEY_DISABLED
- `1008` - API_KEY_PLAN_MINUTE_RATE_LIMIT_REACHED
- `1009` - API_KEY_PLAN_DAILY_RATE_LIMIT_REACHED
- `1010` - API_KEY_PLAN_MONTHLY_RATE_LIMIT_REACHED
- `1011` - IP_RATE_LIMIT_REACHED

## Common Parameters

### Cryptocurrency Identifiers
- `id` - CoinMarketCap ID (e.g., "1" for Bitcoin)
- `symbol` - Cryptocurrency symbol (e.g., "BTC")
- `slug` - URL-friendly name (e.g., "bitcoin")

### Fiat Currency Support
93 fiat currencies supported, including:
- USD (2781), EUR (2790), GBP (2791), CNY (2787), JPY (2797), KRW (2798)

### Pagination
- `start` - Offset (1-based index)
- `limit` - Number of results

### Conversion
- `convert` - Comma-separated list of currencies (e.g., "USD,BTC,EUR")
- `convert_id` - Same as convert but using CMC IDs

## API Endpoints

### Cryptocurrency Endpoints

#### GET /v1/cryptocurrency/map
Returns mapping of all cryptocurrencies to CMC IDs.

**Parameters:**
- `listing_status` (string): "active", "inactive", "untracked"
- `start` (integer): Starting index (default: 1)
- `limit` (integer): Results per page (max: 5000)
- `sort` (string): "id", "cmc_rank"
- `symbol` (string): Filter by symbols
- `aux` (string): "platform,first_historical_data,last_historical_data,is_active,status"

---

#### GET /v2/cryptocurrency/info
Returns metadata for cryptocurrencies.

**Parameters:**
- `id` (string): Comma-separated CMC IDs
- `slug` (string): Comma-separated slugs
- `symbol` (string): Comma-separated symbols
- `address` (string): Contract address
- `aux` (string): "urls,logo,description,tags,platform,date_added,notice,status"

---

#### GET /v1/cryptocurrency/listings/latest
Returns paginated list of all active cryptocurrencies with latest market data.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page (max: 5000)
- `price_min`, `price_max` (number): Price filters
- `market_cap_min`, `market_cap_max` (number): Market cap filters
- `volume_24h_min`, `volume_24h_max` (number): Volume filters
- `circulating_supply_min`, `circulating_supply_max` (number): Supply filters
- `percent_change_24h_min`, `percent_change_24h_max` (number): Change filters
- `convert` (string): Currency conversions
- `convert_id` (string): Currency conversions by ID
- `sort` (string): Sort field
- `sort_dir` (string): "asc", "desc"
- `cryptocurrency_type` (string): "all", "coins", "tokens"
- `tag` (string): Filter by tag
- `aux` (string): Supplemental fields

**Sort Options:**
`market_cap`, `market_cap_strict`, `name`, `symbol`, `date_added`, `price`, `circulating_supply`, `total_supply`, `max_supply`, `num_market_pairs`, `volume_24h`, `percent_change_1h`, `percent_change_24h`, `percent_change_7d`, `market_cap_by_total_supply_strict`, `volume_7d`, `volume_30d`

---

#### GET /v1/cryptocurrency/listings/historical
Returns historical daily rankings.

**Parameters:**
- `date` (string): Date to query (required)
- Other parameters same as listings/latest

---

#### GET /v2/cryptocurrency/quotes/latest
Returns latest market quotes for specific cryptocurrencies.

**Parameters:**
- `id` (string): Comma-separated CMC IDs
- `slug` (string): Comma-separated slugs
- `symbol` (string): Comma-separated symbols
- `convert` (string): Currency conversions
- `convert_id` (string): Currency conversions by ID
- `aux` (string): Supplemental fields
- `skip_invalid` (boolean): Skip invalid lookups

---

#### GET /v2/cryptocurrency/quotes/historical
Returns historical quotes.

**Parameters:**
- `id` (string): CMC IDs
- `symbol` (string): Symbols
- `time_start` (string): Start time (Unix or ISO 8601)
- `time_end` (string): End time
- `count` (number): Number of intervals (max: 10000)
- `interval` (string): Time interval
- `convert` (string): Currency conversions
- `aux` (string): Supplemental fields

**Interval Options:**
Calendar: `hourly`, `daily`, `weekly`, `monthly`, `yearly`
Relative: `5m`, `10m`, `15m`, `30m`, `45m`, `1h`, `2h`, `3h`, `4h`, `6h`, `12h`, `24h`, `1d`, `2d`, `3d`, `7d`, `14d`, `15d`, `30d`, `60d`, `90d`, `365d`

---

#### GET /v3/cryptocurrency/quotes/historical
Returns historical quotes (v3 endpoint with updated response format).

**Parameters:**
- Same as v2/cryptocurrency/quotes/historical
- Note: v3 response format may differ from v2

---

#### GET /v2/cryptocurrency/market-pairs/latest
Returns market pairs for a cryptocurrency.

**Parameters:**
- `id` (string): CMC ID
- `slug` (string): Slug
- `symbol` (string): Symbol
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `aux` (string): Supplemental fields
- `matched_id` (string): Filter by paired currency ID
- `matched_symbol` (string): Filter by paired currency symbol
- `category` (string): "all", "spot", "derivatives", "otc", "perpetual"
- `fee_type` (string): "all", "percentage", "no-fees", "transactional-mining", "unknown"
- `convert` (string): Currency conversions

---

#### GET /v2/cryptocurrency/ohlcv/latest
Returns latest OHLCV data for current UTC day.

**Parameters:**
- `id` (string): CMC IDs
- `symbol` (string): Symbols
- `convert` (string): Currency conversions
- `convert_id` (string): Currency conversions by ID
- `skip_invalid` (boolean): Skip invalid lookups

---

#### GET /v2/cryptocurrency/ohlcv/historical
Returns historical OHLCV data.

**Parameters:**
- `id` (string): CMC IDs
- `slug` (string): Slugs
- `symbol` (string): Symbols
- `time_period` (string): "daily", "hourly"
- `time_start` (string): Start time
- `time_end` (string): End time
- `count` (number): Number of periods
- `interval` (string): Sampling interval
- `convert` (string): Currency conversions

---

#### GET /v2/cryptocurrency/price-performance-stats/latest
Returns price performance statistics.

**Parameters:**
- `id` (string): CMC IDs
- `slug` (string): Slugs
- `symbol` (string): Symbols
- `time_period` (string): "all_time", "yesterday", "24h", "7d", "30d", "90d", "365d"
- `convert` (string): Currency conversions

---

#### GET /v1/cryptocurrency/categories
Returns all coin categories.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `id` (string): Filter by cryptocurrency IDs
- `slug` (string): Filter by slugs
- `symbol` (string): Filter by symbols

---

#### GET /v1/cryptocurrency/category
Returns single category details.

**Parameters:**
- `id` (string): Category ID (required)
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `convert` (string): Currency conversions

---

#### GET /v1/cryptocurrency/airdrops
Returns list of airdrops.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `status` (string): "ENDED", "ONGOING", "UPCOMING"
- `id` (string): Filter by cryptocurrency ID
- `slug` (string): Filter by slug
- `symbol` (string): Filter by symbol

---

#### GET /v1/cryptocurrency/airdrop
Returns single airdrop details.

**Parameters:**
- `id` (string): Airdrop ID (required)

---

#### GET /v1/cryptocurrency/trending/latest
Returns trending cryptocurrencies by search volume.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `time_period` (string): "24h", "30d", "7d"
- `convert` (string): Currency conversions

---

#### GET /v1/cryptocurrency/trending/most-visited
Returns most visited cryptocurrencies.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `time_period` (string): "24h", "30d", "7d"
- `convert` (string): Currency conversions

---

#### GET /v1/cryptocurrency/trending/gainers-losers
Returns biggest gainers and losers.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `time_period` (string): "1h", "24h", "30d", "7d"
- `convert` (string): Currency conversions
- `sort` (string): "percent_change_24h"
- `sort_dir` (string): "asc", "desc"

---

#### GET /v1/cryptocurrency/listings/new
Returns most recently added cryptocurrencies.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page (max: 5000)
- `convert` (string): Currency conversions
- `convert_id` (string): Currency conversions by ID
- `sort_dir` (string): "asc", "desc"

### Exchange Endpoints

#### GET /v1/exchange/map
Returns mapping of all exchanges to CMC IDs.

**Parameters:**
- `listing_status` (string): "active", "inactive", "untracked"
- `slug` (string): Filter by slugs
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `sort` (string): "volume_24h", "id"
- `aux` (string): Supplemental fields
- `crypto_id` (string): Filter by cryptocurrency

---

#### GET /v1/exchange/info
Returns exchange metadata.

**Parameters:**
- `id` (string): Comma-separated exchange IDs
- `slug` (string): Comma-separated slugs
- `aux` (string): "urls,logo,description,date_launched,notice,status"

---

#### GET /v1/exchange/listings/latest
Returns paginated list of all exchanges.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `sort` (string): "name", "volume_24h", "volume_24h_adjusted", "exchange_score"
- `sort_dir` (string): "asc", "desc"
- `market_type` (string): "fees", "no_fees", "all"
- `category` (string): "all", "spot", "derivatives", "dex", "lending"
- `aux` (string): Supplemental fields
- `convert` (string): Currency conversions

---

#### GET /v1/exchange/quotes/latest
Returns latest aggregate market data for exchanges.

**Parameters:**
- `id` (string): Exchange IDs
- `slug` (string): Slugs
- `convert` (string): Currency conversions
- `aux` (string): Supplemental fields

---

#### GET /v1/exchange/quotes/historical
Returns historical exchange quotes.

**Parameters:**
- `id` (string): Exchange IDs
- `slug` (string): Slugs
- `time_start` (string): Start time
- `time_end` (string): End time
- `count` (number): Number of intervals
- `interval` (string): Time interval
- `convert` (string): Currency conversions

---

#### GET /v1/exchange/market-pairs/latest
Returns active market pairs for an exchange.

**Parameters:**
- `id` (string): Exchange ID
- `slug` (string): Exchange slug
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `aux` (string): Supplemental fields
- `matched_id` (string): Filter by currency ID
- `matched_symbol` (string): Filter by currency symbol
- `category` (string): Market category
- `fee_type` (string): Fee type
- `convert` (string): Currency conversions

---

#### GET /v1/exchange/assets
Returns exchange wallet holdings.

**Parameters:**
- `id` (string): Exchange ID (required)

### Global Metrics Endpoints

#### GET /v1/global-metrics/quotes/latest
Returns latest global cryptocurrency market metrics.

**Parameters:**
- `convert` (string): Currency conversions
- `convert_id` (string): Currency conversions by ID

---

#### GET /v1/global-metrics/quotes/historical
Returns historical global market metrics.

**Parameters:**
- `time_start` (string): Start time
- `time_end` (string): End time
- `count` (number): Number of intervals
- `interval` (string): Time interval
- `convert` (string): Currency conversions
- `aux` (string): Supplemental fields

### DexScan Endpoints (Soft Launch)

#### GET /v4/dex/listings/info
Returns DEX metadata.

**Parameters:**
- `id` (string): DEX IDs
- `aux` (string): Supplemental fields

---

#### GET /v4/dex/listings/quotes
Returns list of all DEXes with market data.

**Parameters:**
- `start` (string): Starting index
- `limit` (string): Results per page
- `sort` (string): Sort field
- `sort_dir` (string): Sort direction
- `type` (string): DEX type
- `aux` (string): Supplemental fields
- `convert_id` (string): Currency conversions

---

#### GET /v4/dex/networks/list
Returns all networks.

**Parameters:**
- `start` (string): Starting index
- `limit` (string): Results per page
- `sort` (string): Sort field
- `sort_dir` (string): Sort direction
- `aux` (string): Supplemental fields

---

#### GET /v4/dex/spot-pairs/latest
Returns latest DEX pair listings.

**Parameters:**
- `network_id` (string): Network IDs
- `network_slug` (string): Network slugs
- `dex_id` (string): DEX IDs
- `dex_slug` (string): DEX slugs
- `base_asset_id` (string): Base asset IDs
- `quote_asset_id` (string): Quote asset IDs
- `scroll_id` (string): Pagination cursor
- `limit` (string): Results per page
- `sort` (string): Sort field
- `sort_dir` (string): Sort direction
- `aux` (string): Supplemental fields
- `convert_id` (string): Currency conversions

---

#### GET /v4/dex/pairs/quotes/latest
Returns latest market quotes for DEX pairs.

**Parameters:**
- `contract_address` (string): Contract addresses
- `network_id` (string): Network IDs
- `network_slug` (string): Network slugs
- `aux` (string): Supplemental fields
- `convert_id` (string): Currency conversions
- `skip_invalid` (string): Skip invalid lookups
- `reverse_order` (string): Reverse pair order

---

#### GET /v4/dex/pairs/ohlcv/latest
Returns latest OHLCV for DEX pairs.

**Parameters:**
- Same as pairs/quotes/latest

---

#### GET /v4/dex/pairs/ohlcv/historical
Returns historical OHLCV for DEX pairs.

**Parameters:**
- `contract_address` (string): Contract address
- `network_id` (string): Network ID
- `network_slug` (string): Network slug
- `time_period` (string): Time period
- `time_start` (string): Start time
- `time_end` (string): End time
- `count` (string): Number of periods
- `interval` (string): Time interval
- `aux` (string): Supplemental fields
- `convert_id` (string): Currency conversions

---

#### GET /v4/dex/pairs/trade/latest
Returns latest 100 trades for a DEX pair.

**Parameters:**
- `contract_address` (string): Contract address
- `network_id` (string): Network ID
- `network_slug` (string): Network slug
- `aux` (string): Supplemental fields
- `convert_id` (string): Currency conversions

### Index Endpoints

#### GET /v3/index/cmc100-latest
Returns latest CoinMarketCap 100 Index value.

**No parameters**

---

#### GET /v3/index/cmc100-historical
Returns historical CMC 100 Index values.

**Parameters:**
- `time_start` (string): Start time
- `time_end` (string): End time
- `count` (string): Number of periods
- `interval` (string): "5m", "15m", "daily"

### Fear and Greed Index Endpoints

#### GET /v3/fear-and-greed/latest
Returns latest CMC Crypto Fear and Greed value.

**No parameters**

---

#### GET /v3/fear-and-greed/historical
Returns historical Fear and Greed values.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page (max: 500)

### Fiat Endpoints

#### GET /v1/fiat/map
Returns mapping of all supported fiat currencies.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `sort` (string): "name", "id"
- `include_metals` (boolean): Include precious metals

### Tools Endpoints

#### GET /v2/tools/price-conversion
Converts amounts between currencies.

**Parameters:**
- `amount` (number): Amount to convert (required)
- `id` (string): Source currency CMC ID
- `symbol` (string): Source currency symbol
- `time` (string): Historical timestamp
- `convert` (string): Target currencies
- `convert_id` (string): Target currency IDs

---

#### GET /v1/tools/postman
Returns Postman collection for API.

**No parameters**

### Blockchain Endpoints

#### GET /v1/blockchain/statistics/latest
Returns blockchain statistics (BTC, LTC, ETH supported).

**Parameters:**
- `id` (string): CMC IDs
- `symbol` (string): Symbols
- `slug` (string): Slugs

### Content Endpoints

#### GET /v1/content/latest
Returns cryptocurrency-related news and Alexandria articles.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `category` (string): Content category filter
- `cryptocurrency_id` (string): Filter by cryptocurrency ID
- `language` (string): Language filter
- `sort` (string): Sort field

---

#### GET /v1/content/posts/top
Returns top cryptocurrency-related posts.

**Parameters:**
- `time_period` (string): Time period filter
- `cryptocurrency_id` (string): Filter by cryptocurrency ID
- `start` (integer): Starting index
- `limit` (integer): Results per page

---

#### GET /v1/content/posts/latest
Returns latest cryptocurrency-related posts.

**Parameters:**
- `cryptocurrency_id` (string): Filter by cryptocurrency ID
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `sort` (string): Sort field

---

#### GET /v1/content/posts/comments
Returns comments for a specific post.

**Parameters:**
- `post_id` (string): Post ID (required)
- `start` (integer): Starting index
- `limit` (integer): Results per page

### Community Endpoints

#### GET /v1/community/trending/topic
Returns trending community topics.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `time_period` (string): Time period filter

---

#### GET /v1/community/trending/token
Returns trending community tokens.

**Parameters:**
- `start` (integer): Starting index
- `limit` (integer): Results per page
- `time_period` (string): Time period filter

### Key Management Endpoints

#### GET /v1/key/info
Returns API key details and usage stats.

**No parameters**

## Rate Limits and Credit Usage

### Rate Limits
- Basic: 30 calls/minute
- Hobbyist: 30 calls/minute
- Startup: 60 calls/minute
- Standard: 60 calls/minute
- Professional: 90 calls/minute
- Enterprise: 120 calls/minute

### Credit Usage
- Most endpoints: 1 credit per call
- Additional credits for:
  - Each 100-200 data points returned (varies by endpoint)
  - Each currency conversion beyond the first
  - Bundled requests (100 resources = 1 credit)

### Best Practices

1. **Use CMC IDs instead of symbols** - Symbols are not unique
2. **Implement caching** - Cache frequently accessed data
3. **Handle rate limits** - Implement exponential backoff
4. **Parse responses properly** - Use JSON parsing, not regex
5. **Validate fields** - Add robust field validation
6. **Use appropriate endpoints** - Use listing endpoints for all coins, quote endpoints for specific coins

## Notes

- All timestamps are in UTC
- Date formats support Unix timestamps and ISO 8601
- Maximum bundle size is typically 100-120 items per request
- Historical data availability varies by plan
- Sandbox environment uses test API key: `b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c`

## Complete Endpoint List

### Cryptocurrency (19 endpoints)
1. GET /v1/cryptocurrency/map
2. GET /v2/cryptocurrency/info
3. GET /v1/cryptocurrency/listings/latest
4. GET /v1/cryptocurrency/listings/historical
5. GET /v1/cryptocurrency/listings/new
6. GET /v2/cryptocurrency/quotes/latest
7. GET /v2/cryptocurrency/quotes/historical
8. GET /v3/cryptocurrency/quotes/historical
9. GET /v2/cryptocurrency/market-pairs/latest
10. GET /v2/cryptocurrency/ohlcv/latest
11. GET /v2/cryptocurrency/ohlcv/historical
12. GET /v2/cryptocurrency/price-performance-stats/latest
13. GET /v1/cryptocurrency/categories
14. GET /v1/cryptocurrency/category
15. GET /v1/cryptocurrency/airdrops
16. GET /v1/cryptocurrency/airdrop
17. GET /v1/cryptocurrency/trending/latest
18. GET /v1/cryptocurrency/trending/most-visited
19. GET /v1/cryptocurrency/trending/gainers-losers

### DexScan (8 endpoints)
1. GET /v4/dex/listings/info
2. GET /v4/dex/listings/quotes
3. GET /v4/dex/networks/list
4. GET /v4/dex/spot-pairs/latest
5. GET /v4/dex/pairs/quotes/latest
6. GET /v4/dex/pairs/ohlcv/latest
7. GET /v4/dex/pairs/ohlcv/historical
8. GET /v4/dex/pairs/trade/latest

### Exchange (7 endpoints)
1. GET /v1/exchange/map
2. GET /v1/exchange/info
3. GET /v1/exchange/listings/latest
4. GET /v1/exchange/quotes/latest
5. GET /v1/exchange/quotes/historical
6. GET /v1/exchange/market-pairs/latest
7. GET /v1/exchange/assets

### Global Metrics (2 endpoints)
1. GET /v1/global-metrics/quotes/latest
2. GET /v1/global-metrics/quotes/historical

### Index (2 endpoints)
1. GET /v3/index/cmc100-latest
2. GET /v3/index/cmc100-historical

### Fear and Greed (2 endpoints)
1. GET /v3/fear-and-greed/latest
2. GET /v3/fear-and-greed/historical

### Tools (2 endpoints)
1. GET /v2/tools/price-conversion
2. GET /v1/tools/postman

### Content (4 endpoints)
1. GET /v1/content/latest
2. GET /v1/content/posts/top
3. GET /v1/content/posts/latest
4. GET /v1/content/posts/comments

### Community (2 endpoints)
1. GET /v1/community/trending/topic
2. GET /v1/community/trending/token

### Other (3 endpoints)
1. GET /v1/fiat/map
2. GET /v1/blockchain/statistics/latest
3. GET /v1/key/info

**Total: 53 endpoints**
