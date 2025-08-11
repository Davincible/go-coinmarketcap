// Package coinmarketcap provides types and structures for CoinMarketCap API responses.
package coinmarketcap

import "time"

// APIResponse represents the standard response format from CoinMarketCap API.
// All API endpoints return data in this consistent structure with generic data type T.
type APIResponse[T any] struct {
	Data   T      `json:"data"`
	Status Status `json:"status"`
}

// Status contains metadata about the API response including error information and credit usage.
type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage *string   `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}

// Quote represents price and market data for a cryptocurrency in a specific currency.
type Quote struct {
	Price                 *float64   `json:"price"`
	Volume24h             *float64   `json:"volume_24h"`
	VolumeChange24h       *float64   `json:"volume_change_24h"`
	PercentChange1h       *float64   `json:"percent_change_1h"`
	PercentChange24h      *float64   `json:"percent_change_24h"`
	PercentChange7d       *float64   `json:"percent_change_7d"`
	PercentChange30d      *float64   `json:"percent_change_30d"`
	PercentChange60d      *float64   `json:"percent_change_60d"`
	PercentChange90d      *float64   `json:"percent_change_90d"`
	MarketCap             *float64   `json:"market_cap"`
	MarketCapDominance    *float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap *float64   `json:"fully_diluted_market_cap"`
	TVL                   *float64   `json:"tvl"`
	LastUpdated           *time.Time `json:"last_updated"`
}

// CryptocurrencyMap represents basic cryptocurrency mapping information.
type CryptocurrencyMap struct {
	ID                  int        `json:"id"`
	Name                string     `json:"name"`
	Symbol              string     `json:"symbol"`
	Slug                string     `json:"slug"`
	IsActive            *int       `json:"is_active,omitempty"`
	Status              *int       `json:"status,omitempty"`
	FirstHistoricalData *time.Time `json:"first_historical_data,omitempty"`
	LastHistoricalData  *time.Time `json:"last_historical_data,omitempty"`
	Platform            *Platform  `json:"platform,omitempty"`
}

// Platform represents blockchain platform information for tokens.
type Platform struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"token_address"`
}

// CryptocurrencyInfo contains detailed metadata about a cryptocurrency.
type CryptocurrencyInfo struct {
	ID                            int                 `json:"id"`
	Name                          string              `json:"name"`
	Symbol                        string              `json:"symbol"`
	Category                      string              `json:"category"`
	Description                   *string             `json:"description"`
	Slug                          string              `json:"slug"`
	Logo                          string              `json:"logo"`
	Subreddit                     *string             `json:"subreddit"`
	Notice                        *string             `json:"notice"`
	Tags                          []string            `json:"tags"`
	TagNames                      []string            `json:"tag-names"`
	TagGroups                     []string            `json:"tag-groups"`
	URLs                          map[string][]string `json:"urls"`
	Platform                      *Platform           `json:"platform"`
	DateAdded                     time.Time           `json:"date_added"`
	TwitterUsername               *string             `json:"twitter_username"`
	IsHidden                      *int                `json:"is_hidden"`
	DateLaunched                  *time.Time          `json:"date_launched"`
	ContractAddress               []ContractAddress   `json:"contract_address"`
	SelfReportedCirculatingSupply *float64            `json:"self_reported_circulating_supply"`
	SelfReportedTags              []string            `json:"self_reported_tags"`
	SelfReportedMarketCap         *float64            `json:"self_reported_market_cap"`
	InfiniteSupply                *bool               `json:"infinite_supply"`
}

// ContractAddress represents a smart contract address and its platform.
type ContractAddress struct {
	ContractAddress string   `json:"contract_address"`
	Platform        Platform `json:"platform"`
}

// CryptocurrencyListing represents a cryptocurrency with market data in listings.
type CryptocurrencyListing struct {
	ID                            int               `json:"id"`
	Name                          string            `json:"name"`
	Symbol                        string            `json:"symbol"`
	Slug                          string            `json:"slug"`
	NumMarketPairs                *int              `json:"num_market_pairs"`
	DateAdded                     time.Time         `json:"date_added"`
	Tags                          []string          `json:"tags"`
	MaxSupply                     *float64          `json:"max_supply"`
	CirculatingSupply             *float64          `json:"circulating_supply"`
	TotalSupply                   *float64          `json:"total_supply"`
	InfiniteSupply                *bool             `json:"infinite_supply"`
	Platform                      *Platform         `json:"platform"`
	CMCRank                       *int              `json:"cmc_rank"`
	SelfReportedCirculatingSupply *float64          `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *float64          `json:"self_reported_market_cap"`
	TVLRatio                      *float64          `json:"tvl_ratio"`
	LastUpdated                   time.Time         `json:"last_updated"`
	Quote                         map[string]*Quote `json:"quote"`
}

// CryptocurrencyQuote represents current market data for a cryptocurrency.
type CryptocurrencyQuote struct {
	ID                            int               `json:"id"`
	Name                          string            `json:"name"`
	Symbol                        string            `json:"symbol"`
	Slug                          string            `json:"slug"`
	IsActive                      *int              `json:"is_active"`
	IsFiat                        *int              `json:"is_fiat"`
	NumMarketPairs                *int              `json:"num_market_pairs"`
	DateAdded                     time.Time         `json:"date_added"`
	Tags                          []string          `json:"tags"`
	MaxSupply                     *float64          `json:"max_supply"`
	CirculatingSupply             *float64          `json:"circulating_supply"`
	TotalSupply                   *float64          `json:"total_supply"`
	InfiniteSupply                *bool             `json:"infinite_supply"`
	Platform                      *Platform         `json:"platform"`
	CMCRank                       *int              `json:"cmc_rank"`
	SelfReportedCirculatingSupply *float64          `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *float64          `json:"self_reported_market_cap"`
	TVLRatio                      *float64          `json:"tvl_ratio"`
	LastUpdated                   time.Time         `json:"last_updated"`
	Quote                         map[string]*Quote `json:"quote"`
}

// HistoricalQuote represents historical price data at a specific timestamp.
type HistoricalQuote struct {
	Timestamp      time.Time         `json:"timestamp"`
	Quote          map[string]*Quote `json:"quote"`
	SearchInterval *string           `json:"search_interval,omitempty"`
}

// MarketPair represents a trading pair on an exchange with market data.
type MarketPair struct {
	ExchangeID       int                `json:"exchange_id"`
	ExchangeName     string             `json:"exchange_name"`
	ExchangeSlug     string             `json:"exchange_slug"`
	ExchangeNotice   *string            `json:"exchange_notice"`
	MarketID         int                `json:"market_id"`
	MarketPair       string             `json:"market_pair"`
	MarketPairBase   MarketPairCurrency `json:"market_pair_base"`
	MarketPairQuote  MarketPairCurrency `json:"market_pair_quote"`
	MarketURL        *string            `json:"market_url"`
	MarketScore      *float64           `json:"market_score"`
	MarketReputation *float64           `json:"market_reputation"`
	Category         string             `json:"category"`
	FeeType          string             `json:"fee_type"`
	OutlierDetected  *int               `json:"outlier_detected"`
	ExcludedVolume   *float64           `json:"excluded_volume"`
	Quote            map[string]*Quote  `json:"quote"`
	LastUpdated      time.Time          `json:"last_updated"`
}

// MarketPairCurrency represents currency information within a market pair.
type MarketPairCurrency struct {
	CurrencyID     int    `json:"currency_id"`
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	CurrencySlug   string `json:"currency_slug"`
	ExchangeSymbol string `json:"exchange_symbol"`
}

// OHLCV represents Open, High, Low, Close, Volume data for a time period.
type OHLCV struct {
	TimeOpen  *time.Time        `json:"time_open"`
	TimeClose *time.Time        `json:"time_close"`
	TimeHigh  *time.Time        `json:"time_high"`
	TimeLow   *time.Time        `json:"time_low"`
	Open      *float64          `json:"open"`
	High      *float64          `json:"high"`
	Low       *float64          `json:"low"`
	Close     *float64          `json:"close"`
	Volume    *float64          `json:"volume"`
	MarketCap *float64          `json:"market_cap"`
	Timestamp *time.Time        `json:"timestamp"`
	Quote     map[string]*Quote `json:"quote,omitempty"`
}

// PricePerformanceStats contains ROI performance data for different time periods.
type PricePerformanceStats struct {
	ROI map[string]*PerformancePeriod `json:"roi"`
}

// PerformancePeriod represents performance metrics for a specific time period.
type PerformancePeriod struct {
	Period     string     `json:"period"`
	OpenPrice  *float64   `json:"open_price"`
	HighPrice  *float64   `json:"high_price"`
	LowPrice   *float64   `json:"low_price"`
	ClosePrice *float64   `json:"close_price"`
	OpenTime   *time.Time `json:"open_time"`
	HighTime   *time.Time `json:"high_time"`
	LowTime    *time.Time `json:"low_time"`
	CloseTime  *time.Time `json:"close_time"`
}

// Category represents a cryptocurrency category with aggregate statistics.
type Category struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	NumTokens       int       `json:"num_tokens"`
	AvgPriceChange  *float64  `json:"avg_price_change"`
	MarketCap       *float64  `json:"market_cap"`
	MarketCapChange *float64  `json:"market_cap_change"`
	Volume          *float64  `json:"volume"`
	VolumeChange    *float64  `json:"volume_change"`
	LastUpdated     time.Time `json:"last_updated"`
}

// CategoryDetail extends Category with individual cryptocurrency listings.
type CategoryDetail struct {
	Category
	Coins []CryptocurrencyListing `json:"coins"`
}

// Airdrop represents information about a cryptocurrency airdrop event.
type Airdrop struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	Website          *string    `json:"website"`
	DateAdded        time.Time  `json:"date_added"`
	Status           string     `json:"status"`
	DateStart        *time.Time `json:"date_start"`
	DateEnd          *time.Time `json:"date_end"`
	CryptocurrencyID int        `json:"cryptocurrency_id"`
	Symbol           string     `json:"symbol"`
	Slug             string     `json:"slug"`
}

// Trending represents trending cryptocurrency data based on search volume or other metrics.
type Trending struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Symbol      string            `json:"symbol"`
	Slug        string            `json:"slug"`
	CMCRank     *int              `json:"cmc_rank"`
	SearchScore *float64          `json:"search_score"`
	LastUpdated time.Time         `json:"last_updated"`
	Quote       map[string]*Quote `json:"quote,omitempty"`
}

type ExchangeMap struct {
	ID                  int        `json:"id"`
	Name                string     `json:"name"`
	Slug                string     `json:"slug"`
	IsActive            *int       `json:"is_active,omitempty"`
	Status              *string    `json:"status,omitempty"`
	FirstHistoricalData *time.Time `json:"first_historical_data,omitempty"`
	LastHistoricalData  *time.Time `json:"last_historical_data,omitempty"`
}

type ExchangeInfo struct {
	ID             int                 `json:"id"`
	Name           string              `json:"name"`
	Slug           string              `json:"slug"`
	Logo           string              `json:"logo"`
	Description    *string             `json:"description"`
	DateLaunched   *time.Time          `json:"date_launched"`
	Notice         *string             `json:"notice"`
	Countries      []string            `json:"countries"`
	Fiats          []string            `json:"fiats"`
	Tags           []string            `json:"tags"`
	Type           *string             `json:"type"`
	IsHidden       *int                `json:"is_hidden"`
	IsDistributed  *int                `json:"is_distributed"`
	MakerFee       *float64            `json:"maker_fee"`
	TakerFee       *float64            `json:"taker_fee"`
	WeeklyVisits   *int                `json:"weekly_visits"`
	SpotVolumeUsd  *float64            `json:"spot_volume_usd"`
	SpotVolumeRank *int                `json:"spot_volume_rank"`
	URLs           map[string][]string `json:"urls"`
}

type ExchangeListing struct {
	ID                     int               `json:"id"`
	Name                   string            `json:"name"`
	Slug                   string            `json:"slug"`
	NumMarketPairs         *int              `json:"num_market_pairs"`
	NumCoins               *int              `json:"num_coins"`
	NumCryptocurrencies    *int              `json:"num_cryptocurrencies"`
	NumFiats               *int              `json:"num_fiats"`
	DateLaunched           *time.Time        `json:"date_launched"`
	Fiats                  []string          `json:"fiats"`
	Tags                   []string          `json:"tags"`
	Type                   *string           `json:"type"`
	ExchangeScore          *float64          `json:"exchange_score"`
	Derivatives            *float64          `json:"derivatives"`
	WeeklyVisits           *int              `json:"weekly_visits"`
	SpotVolumeUsd          *float64          `json:"spot_volume_usd"`
	DerivativesVolumeUsd   *float64          `json:"derivatives_volume_usd"`
	Volume24hReported      *float64          `json:"volume_24h_reported"`
	Volume24hAdjusted      *float64          `json:"volume_24h_adjusted"`
	Volume7dReported       *float64          `json:"volume_7d_reported"`
	Volume30dReported      *float64          `json:"volume_30d_reported"`
	PercentChangeVolume24h *float64          `json:"percent_change_volume_24h"`
	PercentChangeVolume7d  *float64          `json:"percent_change_volume_7d"`
	PercentChangeVolume30d *float64          `json:"percent_change_volume_30d"`
	TrafficScore           *float64          `json:"traffic_score"`
	LiquidityScore         *float64          `json:"liquidity_score"`
	LastUpdated            time.Time         `json:"last_updated"`
	Quote                  map[string]*Quote `json:"quote,omitempty"`
}

type ExchangeQuote struct {
	ID                     int               `json:"id"`
	Name                   string            `json:"name"`
	Slug                   string            `json:"slug"`
	NumMarketPairs         *int              `json:"num_market_pairs"`
	Volume24hReported      *float64          `json:"volume_24h_reported"`
	Volume24hAdjusted      *float64          `json:"volume_24h_adjusted"`
	Volume7dReported       *float64          `json:"volume_7d_reported"`
	Volume30dReported      *float64          `json:"volume_30d_reported"`
	PercentChangeVolume24h *float64          `json:"percent_change_volume_24h"`
	PercentChangeVolume7d  *float64          `json:"percent_change_volume_7d"`
	PercentChangeVolume30d *float64          `json:"percent_change_volume_30d"`
	LastUpdated            time.Time         `json:"last_updated"`
	Quote                  map[string]*Quote `json:"quote,omitempty"`
}

type GlobalMetrics struct {
	BtcDominance                   *float64          `json:"btc_dominance"`
	EthDominance                   *float64          `json:"eth_dominance"`
	IcpDominance                   *float64          `json:"icp_dominance"`
	ActiveCryptocurrencies         *int              `json:"active_cryptocurrencies"`
	TotalCryptocurrencies          *int              `json:"total_cryptocurrencies"`
	ActiveExchanges                *int              `json:"active_exchanges"`
	TotalExchanges                 *int              `json:"total_exchanges"`
	ActiveMarketPairs              *int              `json:"active_market_pairs"`
	TotalMarketPairs               *int              `json:"total_market_pairs"`
	DefiMarketCap                  *float64          `json:"defi_market_cap"`
	DefiMarketCapDominance         *float64          `json:"defi_market_cap_dominance"`
	Defi24hPercentageChange        *float64          `json:"defi_24h_percentage_change"`
	StablecoinMarketCap            *float64          `json:"stablecoin_market_cap"`
	StablecoinMarketCapDominance   *float64          `json:"stablecoin_market_cap_dominance"`
	Stablecoin24hPercentageChange  *float64          `json:"stablecoin_24h_percentage_change"`
	DerivativesMarketCap           *float64          `json:"derivatives_market_cap"`
	DerivativesMarketCapDominance  *float64          `json:"derivatives_market_cap_dominance"`
	Derivatives24hPercentageChange *float64          `json:"derivatives_24h_percentage_change"`
	LastUpdated                    time.Time         `json:"last_updated"`
	Quote                          map[string]*Quote `json:"quote"`
}

type FiatMap struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Symbol string `json:"symbol"`
}

type PriceConversion struct {
	Symbol      string                      `json:"symbol"`
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Amount      float64                     `json:"amount"`
	LastUpdated time.Time                   `json:"last_updated"`
	Quote       map[string]*ConversionQuote `json:"quote"`
}

type ConversionQuote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

type BlockchainStats struct {
	ID                     int        `json:"id"`
	Symbol                 string     `json:"symbol"`
	Name                   string     `json:"name"`
	TotalSupply            string     `json:"total_supply"`
	Beta                   *float64   `json:"beta"`
	CorrelationPearson     *float64   `json:"correlation_pearson"`
	Count24hInterval       *int       `json:"count_24h_interval"`
	Count30dInterval       *int       `json:"count_30d_interval"`
	CountYtdInterval       *int       `json:"count_ytd_interval"`
	FirstBlockTimestamp    time.Time  `json:"first_block_timestamp"`
	FirstPricedTimestamp   time.Time  `json:"first_priced_timestamp"`
	HashAlgorithm          *string    `json:"hash_algorithm"`
	HashrateEma            *string    `json:"hashrate_ema"`
	High24h                *float64   `json:"high_24h"`
	InflationRate          *float64   `json:"inflation_rate"`
	IssueRate              *float64   `json:"issue_rate"`
	LastBlockHeight        *int       `json:"last_block_height"`
	LastBlockTimestamp     time.Time  `json:"last_block_timestamp"`
	LastKnownHashrate      *string    `json:"last_known_hashrate"`
	Low24h                 *float64   `json:"low_24h"`
	MeanBlockTime          *int       `json:"mean_block_time"`
	MeanTxFee              *float64   `json:"mean_tx_fee"`
	MeanTxValue            *float64   `json:"mean_tx_value"`
	MedianTxFee            *float64   `json:"median_tx_fee"`
	MedianTxValue          *float64   `json:"median_tx_value"`
	NextHalvingDate        *time.Time `json:"next_halving_date"`
	NextDifficultyRetarget *time.Time `json:"next_difficulty_retarget"`
	PendingTransactions    *int       `json:"pending_transactions"`
	RewardsEma             *string    `json:"rewards_ema"`
	Sum24hFees             *float64   `json:"sum_24h_fees"`
	Sum24hRewards          *float64   `json:"sum_24h_rewards"`
	Sum24hTransactionCount *int       `json:"sum_24h_transaction_count"`
	Sum24hTxVolume         *string    `json:"sum_24h_tx_volume"`
}

type KeyInfo struct {
	Plan struct {
		Name                             string    `json:"name"`
		CreditLimitDaily                 int       `json:"credit_limit_daily"`
		CreditLimitDailyReset            string    `json:"credit_limit_daily_reset"`
		CreditLimitDailyResetTimestamp   time.Time `json:"credit_limit_daily_reset_timestamp"`
		CreditLimitMonthly               int       `json:"credit_limit_monthly"`
		CreditLimitMonthlyReset          string    `json:"credit_limit_monthly_reset"`
		CreditLimitMonthlyResetTimestamp time.Time `json:"credit_limit_monthly_reset_timestamp"`
		RateLimitMinute                  int       `json:"rate_limit_minute"`
	} `json:"plan"`
	Usage struct {
		CurrentMinute struct {
			RequestsLeft int `json:"requests_left"`
			RequestsMade int `json:"requests_made"`
		} `json:"current_minute"`
		CurrentDay struct {
			CreditsLeft int `json:"credits_left"`
			CreditsUsed int `json:"credits_used"`
		} `json:"current_day"`
		CurrentMonth struct {
			CreditsLeft int `json:"credits_left"`
			CreditsUsed int `json:"credits_used"`
		} `json:"current_month"`
	} `json:"usage"`
}

type ListingSort string

const (
	SortMarketCap              ListingSort = "market_cap"
	SortMarketCapStrict        ListingSort = "market_cap_strict"
	SortName                   ListingSort = "name"
	SortSymbol                 ListingSort = "symbol"
	SortDateAdded              ListingSort = "date_added"
	SortPrice                  ListingSort = "price"
	SortCirculatingSupply      ListingSort = "circulating_supply"
	SortTotalSupply            ListingSort = "total_supply"
	SortMaxSupply              ListingSort = "max_supply"
	SortNumMarketPairs         ListingSort = "num_market_pairs"
	SortVolume24h              ListingSort = "volume_24h"
	SortPercentChange1h        ListingSort = "percent_change_1h"
	SortPercentChange24h       ListingSort = "percent_change_24h"
	SortPercentChange7d        ListingSort = "percent_change_7d"
	SortMarketCapByTotalSupply ListingSort = "market_cap_by_total_supply_strict"
	SortVolume7d               ListingSort = "volume_7d"
	SortVolume30d              ListingSort = "volume_30d"
)

type SortDirection string

const (
	SortAsc  SortDirection = "asc"
	SortDesc SortDirection = "desc"
)

type CryptocurrencyType string

const (
	CryptoTypeAll    CryptocurrencyType = "all"
	CryptoTypeCoins  CryptocurrencyType = "coins"
	CryptoTypeTokens CryptocurrencyType = "tokens"
)

type ListingStatus string

const (
	StatusActive    ListingStatus = "active"
	StatusInactive  ListingStatus = "inactive"
	StatusUntracked ListingStatus = "untracked"
)

type Interval string

const (
	Interval5m      Interval = "5m"
	Interval10m     Interval = "10m"
	Interval15m     Interval = "15m"
	Interval30m     Interval = "30m"
	Interval45m     Interval = "45m"
	Interval1h      Interval = "1h"
	Interval2h      Interval = "2h"
	Interval3h      Interval = "3h"
	Interval4h      Interval = "4h"
	Interval6h      Interval = "6h"
	Interval12h     Interval = "12h"
	Interval24h     Interval = "24h"
	Interval1d      Interval = "1d"
	Interval2d      Interval = "2d"
	Interval3d      Interval = "3d"
	Interval7d      Interval = "7d"
	Interval14d     Interval = "14d"
	Interval15d     Interval = "15d"
	Interval30d     Interval = "30d"
	Interval60d     Interval = "60d"
	Interval90d     Interval = "90d"
	Interval365d    Interval = "365d"
	IntervalHourly  Interval = "hourly"
	IntervalDaily   Interval = "daily"
	IntervalWeekly  Interval = "weekly"
	IntervalMonthly Interval = "monthly"
	IntervalYearly  Interval = "yearly"
)

type TimePeriod string

const (
	TimePeriodAllTime   TimePeriod = "all_time"
	TimePeriodYesterday TimePeriod = "yesterday"
	TimePeriod24h       TimePeriod = "24h"
	TimePeriod7d        TimePeriod = "7d"
	TimePeriod30d       TimePeriod = "30d"
	TimePeriod90d       TimePeriod = "90d"
	TimePeriod365d      TimePeriod = "365d"
	TimePeriod1h        TimePeriod = "1h"
)

type MarketType string

const (
	MarketTypeFees   MarketType = "fees"
	MarketTypeNoFees MarketType = "no_fees"
	MarketTypeAll    MarketType = "all"
)

type ExchangeCategory string

const (
	ExchangeCategoryAll         ExchangeCategory = "all"
	ExchangeCategorySpot        ExchangeCategory = "spot"
	ExchangeCategoryDerivatives ExchangeCategory = "derivatives"
	ExchangeCategoryDex         ExchangeCategory = "dex"
	ExchangeCategoryLending     ExchangeCategory = "lending"
)

type FeeType string

const (
	FeeTypeAll                 FeeType = "all"
	FeeTypePercentage          FeeType = "percentage"
	FeeTypeNoFees              FeeType = "no-fees"
	FeeTypeTransactionalMining FeeType = "transactional-mining"
	FeeTypeUnknown             FeeType = "unknown"
)

type PairCategory string

const (
	PairCategoryAll         PairCategory = "all"
	PairCategorySpot        PairCategory = "spot"
	PairCategoryDerivatives PairCategory = "derivatives"
	PairCategoryOTC         PairCategory = "otc"
	PairCategoryPerpetual   PairCategory = "perpetual"
)

type AirdropStatus string

const (
	AirdropStatusEnded    AirdropStatus = "ENDED"
	AirdropStatusOngoing  AirdropStatus = "ONGOING"
	AirdropStatusUpcoming AirdropStatus = "UPCOMING"
)

type ExchangeSort string

const (
	ExchangeSortName              ExchangeSort = "name"
	ExchangeSortVolume24h         ExchangeSort = "volume_24h"
	ExchangeSortVolume24hAdjusted ExchangeSort = "volume_24h_adjusted"
	ExchangeSortExchangeScore     ExchangeSort = "exchange_score"
	ExchangeSortID                ExchangeSort = "id"
)
