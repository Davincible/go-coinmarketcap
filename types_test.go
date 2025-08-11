package coinmarketcap

import (
	"encoding/json"
	"testing"
)

func TestAPIResponseUnmarshaling(t *testing.T) {
	jsonData := `{
		"data": [
			{
				"id": 1,
				"name": "Bitcoin",
				"symbol": "BTC",
				"slug": "bitcoin"
			}
		],
		"status": {
			"timestamp": "2023-01-01T00:00:00.000Z",
			"error_code": 0,
			"error_message": null,
			"elapsed": 10,
			"credit_count": 1
		}
	}`

	var response APIResponse[[]CryptocurrencyMap]
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(response.Data) != 1 {
		t.Errorf("expected 1 item in data, got %d", len(response.Data))
	}

	crypto := response.Data[0]
	if crypto.ID != 1 {
		t.Errorf("expected ID 1, got %d", crypto.ID)
	}
	if crypto.Name != "Bitcoin" {
		t.Errorf("expected name 'Bitcoin', got %s", crypto.Name)
	}
	if crypto.Symbol != "BTC" {
		t.Errorf("expected symbol 'BTC', got %s", crypto.Symbol)
	}
	if crypto.Slug != "bitcoin" {
		t.Errorf("expected slug 'bitcoin', got %s", crypto.Slug)
	}

	if response.Status.ErrorCode != 0 {
		t.Errorf("expected error code 0, got %d", response.Status.ErrorCode)
	}
	if response.Status.CreditCount != 1 {
		t.Errorf("expected credit count 1, got %d", response.Status.CreditCount)
	}
}

func TestCryptocurrencyListingUnmarshaling(t *testing.T) {
	jsonData := `{
		"id": 1,
		"name": "Bitcoin",
		"symbol": "BTC",
		"slug": "bitcoin",
		"num_market_pairs": 9999,
		"date_added": "2013-04-28T00:00:00.000Z",
		"tags": ["cryptocurrency", "blockchain"],
		"max_supply": 21000000,
		"circulating_supply": 19000000,
		"total_supply": 19000000,
		"infinite_supply": false,
		"platform": null,
		"cmc_rank": 1,
		"self_reported_circulating_supply": null,
		"self_reported_market_cap": null,
		"tvl_ratio": null,
		"last_updated": "2023-01-01T00:00:00.000Z",
		"quote": {
			"USD": {
				"price": 50000.0,
				"volume_24h": 25000000000,
				"volume_change_24h": 5.5,
				"percent_change_1h": 0.5,
				"percent_change_24h": 2.5,
				"percent_change_7d": 10.5,
				"percent_change_30d": 15.5,
				"percent_change_60d": 20.5,
				"percent_change_90d": 25.5,
				"market_cap": 950000000000,
				"market_cap_dominance": 45.5,
				"fully_diluted_market_cap": 1050000000000,
				"tvl": null,
				"last_updated": "2023-01-01T00:00:00.000Z"
			}
		}
	}`

	var listing CryptocurrencyListing
	err := json.Unmarshal([]byte(jsonData), &listing)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if listing.ID != 1 {
		t.Errorf("expected ID 1, got %d", listing.ID)
	}
	if listing.Name != "Bitcoin" {
		t.Errorf("expected name 'Bitcoin', got %s", listing.Name)
	}
	if listing.Symbol != "BTC" {
		t.Errorf("expected symbol 'BTC', got %s", listing.Symbol)
	}

	if listing.NumMarketPairs == nil || *listing.NumMarketPairs != 9999 {
		t.Errorf("expected num market pairs 9999, got %v", listing.NumMarketPairs)
	}

	if len(listing.Tags) != 2 {
		t.Errorf("expected 2 tags, got %d", len(listing.Tags))
	}

	if listing.MaxSupply == nil || *listing.MaxSupply != 21000000 {
		t.Errorf("expected max supply 21000000, got %v", listing.MaxSupply)
	}

	if listing.InfiniteSupply == nil || *listing.InfiniteSupply != false {
		t.Errorf("expected infinite supply false, got %v", listing.InfiniteSupply)
	}

	if listing.CMCRank == nil || *listing.CMCRank != 1 {
		t.Errorf("expected CMC rank 1, got %v", listing.CMCRank)
	}

	// Test quote data
	usdQuote, exists := listing.Quote["USD"]
	if !exists {
		t.Fatal("expected USD quote to exist")
	}

	if usdQuote.Price == nil || *usdQuote.Price != 50000.0 {
		t.Errorf("expected price 50000.0, got %v", usdQuote.Price)
	}

	if usdQuote.Volume24h == nil || *usdQuote.Volume24h != 25000000000 {
		t.Errorf("expected volume 25000000000, got %v", usdQuote.Volume24h)
	}

	if usdQuote.PercentChange24h == nil || *usdQuote.PercentChange24h != 2.5 {
		t.Errorf("expected percent change 2.5, got %v", usdQuote.PercentChange24h)
	}

	if usdQuote.MarketCap == nil || *usdQuote.MarketCap != 950000000000 {
		t.Errorf("expected market cap 950000000000, got %v", usdQuote.MarketCap)
	}
}

func TestOHLCVUnmarshaling(t *testing.T) {
	jsonData := `{
		"time_open": "2023-01-01T00:00:00.000Z",
		"time_close": "2023-01-01T23:59:59.000Z",
		"time_high": "2023-01-01T12:30:00.000Z",
		"time_low": "2023-01-01T06:15:00.000Z",
		"open": 48000.0,
		"high": 52000.0,
		"low": 47000.0,
		"close": 50000.0,
		"volume": 1000000000,
		"market_cap": 950000000000,
		"timestamp": "2023-01-01T00:00:00.000Z"
	}`

	var ohlcv OHLCV
	err := json.Unmarshal([]byte(jsonData), &ohlcv)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if ohlcv.Open == nil || *ohlcv.Open != 48000.0 {
		t.Errorf("expected open 48000.0, got %v", ohlcv.Open)
	}
	if ohlcv.High == nil || *ohlcv.High != 52000.0 {
		t.Errorf("expected high 52000.0, got %v", ohlcv.High)
	}
	if ohlcv.Low == nil || *ohlcv.Low != 47000.0 {
		t.Errorf("expected low 47000.0, got %v", ohlcv.Low)
	}
	if ohlcv.Close == nil || *ohlcv.Close != 50000.0 {
		t.Errorf("expected close 50000.0, got %v", ohlcv.Close)
	}
	if ohlcv.Volume == nil || *ohlcv.Volume != 1000000000 {
		t.Errorf("expected volume 1000000000, got %v", ohlcv.Volume)
	}

	if ohlcv.TimeOpen == nil {
		t.Error("expected time_open to be set")
	}
	if ohlcv.TimeClose == nil {
		t.Error("expected time_close to be set")
	}
}

func TestMarketPairUnmarshaling(t *testing.T) {
	jsonData := `{
		"exchange_id": 270,
		"exchange_name": "Binance",
		"exchange_slug": "binance",
		"exchange_notice": "",
		"market_id": 1234,
		"market_pair": "BTC/USDT",
		"market_pair_base": {
			"currency_id": 1,
			"currency_name": "Bitcoin",
			"currency_symbol": "BTC",
			"currency_slug": "bitcoin",
			"exchange_symbol": "BTC"
		},
		"market_pair_quote": {
			"currency_id": 825,
			"currency_name": "Tether",
			"currency_symbol": "USDT",
			"currency_slug": "tether",
			"exchange_symbol": "USDT"
		},
		"market_url": "https://www.binance.com/en/trade/BTC_USDT",
		"market_score": 0.98,
		"market_reputation": 0.95,
		"category": "spot",
		"fee_type": "percentage",
		"outlier_detected": 0,
		"excluded_volume": 0,
		"quote": {
			"USD": {
				"price": 50000.0,
				"volume_24h": 1000000000,
				"last_updated": "2023-01-01T00:00:00.000Z"
			}
		},
		"last_updated": "2023-01-01T00:00:00.000Z"
	}`

	var pair MarketPair
	err := json.Unmarshal([]byte(jsonData), &pair)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if pair.ExchangeID != 270 {
		t.Errorf("expected exchange ID 270, got %d", pair.ExchangeID)
	}
	if pair.ExchangeName != "Binance" {
		t.Errorf("expected exchange name 'Binance', got %s", pair.ExchangeName)
	}
	if pair.MarketPair != "BTC/USDT" {
		t.Errorf("expected market pair 'BTC/USDT', got %s", pair.MarketPair)
	}

	if pair.MarketPairBase.CurrencyID != 1 {
		t.Errorf("expected base currency ID 1, got %d", pair.MarketPairBase.CurrencyID)
	}
	if pair.MarketPairBase.CurrencySymbol != "BTC" {
		t.Errorf("expected base currency symbol 'BTC', got %s", pair.MarketPairBase.CurrencySymbol)
	}

	if pair.MarketPairQuote.CurrencyID != 825 {
		t.Errorf("expected quote currency ID 825, got %d", pair.MarketPairQuote.CurrencyID)
	}
	if pair.MarketPairQuote.CurrencySymbol != "USDT" {
		t.Errorf("expected quote currency symbol 'USDT', got %s", pair.MarketPairQuote.CurrencySymbol)
	}

	if pair.MarketScore == nil || *pair.MarketScore != 0.98 {
		t.Errorf("expected market score 0.98, got %v", pair.MarketScore)
	}

	if pair.Category != "spot" {
		t.Errorf("expected category 'spot', got %s", pair.Category)
	}

	// Test quote
	usdQuote, exists := pair.Quote["USD"]
	if !exists {
		t.Fatal("expected USD quote to exist")
	}
	if usdQuote.Price == nil || *usdQuote.Price != 50000.0 {
		t.Errorf("expected price 50000.0, got %v", usdQuote.Price)
	}
}

func TestGlobalMetricsUnmarshaling(t *testing.T) {
	jsonData := `{
		"btc_dominance": 45.5,
		"eth_dominance": 18.2,
		"icp_dominance": null,
		"active_cryptocurrencies": 5000,
		"total_cryptocurrencies": 8000,
		"active_exchanges": 500,
		"total_exchanges": 600,
		"active_market_pairs": 50000,
		"total_market_pairs": 60000,
		"defi_market_cap": 100000000000,
		"defi_market_cap_dominance": 5.5,
		"defi_24h_percentage_change": 2.5,
		"last_updated": "2023-01-01T00:00:00.000Z",
		"quote": {
			"USD": {
				"market_cap": 2000000000000,
				"volume_24h": 100000000000,
				"percent_change_24h": 1.5,
				"last_updated": "2023-01-01T00:00:00.000Z"
			}
		}
	}`

	var metrics GlobalMetrics
	err := json.Unmarshal([]byte(jsonData), &metrics)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if metrics.BtcDominance == nil || *metrics.BtcDominance != 45.5 {
		t.Errorf("expected BTC dominance 45.5, got %v", metrics.BtcDominance)
	}
	if metrics.EthDominance == nil || *metrics.EthDominance != 18.2 {
		t.Errorf("expected ETH dominance 18.2, got %v", metrics.EthDominance)
	}
	if metrics.IcpDominance != nil {
		t.Errorf("expected ICP dominance to be nil, got %v", metrics.IcpDominance)
	}
	if metrics.ActiveCryptocurrencies == nil || *metrics.ActiveCryptocurrencies != 5000 {
		t.Errorf("expected active cryptocurrencies 5000, got %v", metrics.ActiveCryptocurrencies)
	}

	usdQuote, exists := metrics.Quote["USD"]
	if !exists {
		t.Fatal("expected USD quote to exist")
	}
	if usdQuote.MarketCap == nil || *usdQuote.MarketCap != 2000000000000 {
		t.Errorf("expected market cap 2000000000000, got %v", usdQuote.MarketCap)
	}
}

func TestConstantValues(t *testing.T) {
	// Test ListingSort constants
	if SortMarketCap != "market_cap" {
		t.Errorf("expected SortMarketCap to be 'market_cap', got %s", SortMarketCap)
	}
	if SortPrice != "price" {
		t.Errorf("expected SortPrice to be 'price', got %s", SortPrice)
	}

	// Test SortDirection constants
	if SortAsc != "asc" {
		t.Errorf("expected SortAsc to be 'asc', got %s", SortAsc)
	}
	if SortDesc != "desc" {
		t.Errorf("expected SortDesc to be 'desc', got %s", SortDesc)
	}

	// Test CryptocurrencyType constants
	if CryptoTypeAll != "all" {
		t.Errorf("expected CryptoTypeAll to be 'all', got %s", CryptoTypeAll)
	}
	if CryptoTypeCoins != "coins" {
		t.Errorf("expected CryptoTypeCoins to be 'coins', got %s", CryptoTypeCoins)
	}

	// Test Interval constants
	if Interval1h != "1h" {
		t.Errorf("expected Interval1h to be '1h', got %s", Interval1h)
	}
	if IntervalDaily != "daily" {
		t.Errorf("expected IntervalDaily to be 'daily', got %s", IntervalDaily)
	}
}
