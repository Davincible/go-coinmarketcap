package coinmarketcap

import (
	"context"
)

type CryptocurrencyMapOptions struct {
	ListingStatus *ListingStatus
	Start         *int
	Limit         *int
	Sort          *string
	Symbol        []string
	Aux           []string
}

func (c *Client) GetCryptocurrencyMap(ctx context.Context, opts *CryptocurrencyMapOptions) (*APIResponse[[]CryptocurrencyMap], error) {
	params := NewParamBuilder()

	if opts != nil {
		if opts.ListingStatus != nil {
			params.Add("listing_status", string(*opts.ListingStatus))
		}
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.Add("sort", *opts.Sort)
		params.AddStringSlice("symbol", opts.Symbol)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[[]CryptocurrencyMap](c, ctx, "/v1/cryptocurrency/map", &RequestOptions[[]CryptocurrencyMap]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyInfoOptions struct {
	ID      []int
	Slug    []string
	Symbol  []string
	Address []string
	Aux     []string
}

func (c *Client) GetCryptocurrencyInfo(ctx context.Context, opts *CryptocurrencyInfoOptions) (*APIResponse[map[string]CryptocurrencyInfo], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.AddStringSlice("symbol", opts.Symbol)
		params.AddStringSlice("address", opts.Address)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[map[string]CryptocurrencyInfo](c, ctx, "/v2/cryptocurrency/info", &RequestOptions[map[string]CryptocurrencyInfo]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyListingsOptions struct {
	Start                *int
	Limit                *int
	PriceMin             *float64
	PriceMax             *float64
	MarketCapMin         *float64
	MarketCapMax         *float64
	Volume24hMin         *float64
	Volume24hMax         *float64
	CirculatingSupplyMin *float64
	CirculatingSupplyMax *float64
	PercentChange24hMin  *float64
	PercentChange24hMax  *float64
	Convert              []string
	ConvertID            []int
	Sort                 *ListingSort
	SortDir              *SortDirection
	CryptocurrencyType   *CryptocurrencyType
	Tag                  *string
	Aux                  []string
}

func (c *Client) GetCryptocurrencyListingsLatest(ctx context.Context, opts *CryptocurrencyListingsOptions) (*APIResponse[[]CryptocurrencyListing], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.AddFloat("price_min", opts.PriceMin)
		params.AddFloat("price_max", opts.PriceMax)
		params.AddFloat("market_cap_min", opts.MarketCapMin)
		params.AddFloat("market_cap_max", opts.MarketCapMax)
		params.AddFloat("volume_24h_min", opts.Volume24hMin)
		params.AddFloat("volume_24h_max", opts.Volume24hMax)
		params.AddFloat("circulating_supply_min", opts.CirculatingSupplyMin)
		params.AddFloat("circulating_supply_max", opts.CirculatingSupplyMax)
		params.AddFloat("percent_change_24h_min", opts.PercentChange24hMin)
		params.AddFloat("percent_change_24h_max", opts.PercentChange24hMax)
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
		if opts.Sort != nil {
			params.Add("sort", string(*opts.Sort))
		}
		if opts.SortDir != nil {
			params.Add("sort_dir", string(*opts.SortDir))
		}
		if opts.CryptocurrencyType != nil {
			params.Add("cryptocurrency_type", string(*opts.CryptocurrencyType))
		}
		params.Add("tag", *opts.Tag)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[[]CryptocurrencyListing](c, ctx, "/v1/cryptocurrency/listings/latest", &RequestOptions[[]CryptocurrencyListing]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyListingsHistoricalOptions struct {
	Date string
	CryptocurrencyListingsOptions
}

func (c *Client) GetCryptocurrencyListingsHistorical(ctx context.Context, opts *CryptocurrencyListingsHistoricalOptions) (*APIResponse[[]CryptocurrencyListing], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.Add("date", opts.Date)
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.AddFloat("price_min", opts.PriceMin)
		params.AddFloat("price_max", opts.PriceMax)
		params.AddFloat("market_cap_min", opts.MarketCapMin)
		params.AddFloat("market_cap_max", opts.MarketCapMax)
		params.AddFloat("volume_24h_min", opts.Volume24hMin)
		params.AddFloat("volume_24h_max", opts.Volume24hMax)
		params.AddFloat("circulating_supply_min", opts.CirculatingSupplyMin)
		params.AddFloat("circulating_supply_max", opts.CirculatingSupplyMax)
		params.AddFloat("percent_change_24h_min", opts.PercentChange24hMin)
		params.AddFloat("percent_change_24h_max", opts.PercentChange24hMax)
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
		if opts.Sort != nil {
			params.Add("sort", string(*opts.Sort))
		}
		if opts.SortDir != nil {
			params.Add("sort_dir", string(*opts.SortDir))
		}
		if opts.CryptocurrencyType != nil {
			params.Add("cryptocurrency_type", string(*opts.CryptocurrencyType))
		}
		params.Add("tag", *opts.Tag)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[[]CryptocurrencyListing](c, ctx, "/v1/cryptocurrency/listings/historical", &RequestOptions[[]CryptocurrencyListing]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyListingsNewOptions struct {
	Start     *int
	Limit     *int
	Convert   []string
	ConvertID []int
	SortDir   *SortDirection
}

func (c *Client) GetCryptocurrencyListingsNew(ctx context.Context, opts *CryptocurrencyListingsNewOptions) (*APIResponse[[]CryptocurrencyListing], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
		if opts.SortDir != nil {
			params.Add("sort_dir", string(*opts.SortDir))
		}
	}

	return get[[]CryptocurrencyListing](c, ctx, "/v1/cryptocurrency/listings/new", &RequestOptions[[]CryptocurrencyListing]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyQuotesOptions struct {
	ID          []int
	Slug        []string
	Symbol      []string
	Convert     []string
	ConvertID   []int
	Aux         []string
	SkipInvalid *bool
}

func (c *Client) GetCryptocurrencyQuotesLatest(ctx context.Context, opts *CryptocurrencyQuotesOptions) (*APIResponse[map[string]CryptocurrencyQuote], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.AddStringSlice("symbol", opts.Symbol)
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
		params.AddStringSlice("aux", opts.Aux)
		params.AddBool("skip_invalid", opts.SkipInvalid)
	}

	return get[map[string]CryptocurrencyQuote](c, ctx, "/v2/cryptocurrency/quotes/latest", &RequestOptions[map[string]CryptocurrencyQuote]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyQuotesHistoricalOptions struct {
	ID        []int
	Symbol    []string
	TimeStart *string
	TimeEnd   *string
	Count     *int
	Interval  *Interval
	Convert   []string
	ConvertID []int
	Aux       []string
}

func (c *Client) GetCryptocurrencyQuotesHistorical(ctx context.Context, opts *CryptocurrencyQuotesHistoricalOptions) (*APIResponse[map[string][]HistoricalQuote], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("symbol", opts.Symbol)
		params.Add("time_start", *opts.TimeStart)
		params.Add("time_end", *opts.TimeEnd)
		params.AddInt("count", opts.Count)
		if opts.Interval != nil {
			params.Add("interval", string(*opts.Interval))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[map[string][]HistoricalQuote](c, ctx, "/v2/cryptocurrency/quotes/historical", &RequestOptions[map[string][]HistoricalQuote]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetCryptocurrencyQuotesHistoricalV3(ctx context.Context, opts *CryptocurrencyQuotesHistoricalOptions) (*APIResponse[map[string][]HistoricalQuote], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("symbol", opts.Symbol)
		params.Add("time_start", *opts.TimeStart)
		params.Add("time_end", *opts.TimeEnd)
		params.AddInt("count", opts.Count)
		if opts.Interval != nil {
			params.Add("interval", string(*opts.Interval))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[map[string][]HistoricalQuote](c, ctx, "/v3/cryptocurrency/quotes/historical", &RequestOptions[map[string][]HistoricalQuote]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyMarketPairsOptions struct {
	ID            *int
	Slug          *string
	Symbol        *string
	Start         *int
	Limit         *int
	Aux           []string
	MatchedID     []int
	MatchedSymbol []string
	Category      *PairCategory
	FeeType       *FeeType
	Convert       []string
	ConvertID     []int
}

func (c *Client) GetCryptocurrencyMarketPairsLatest(ctx context.Context, opts *CryptocurrencyMarketPairsOptions) (*APIResponse[map[string][]MarketPair], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("id", opts.ID)
		params.Add("slug", *opts.Slug)
		params.Add("symbol", *opts.Symbol)
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.AddStringSlice("aux", opts.Aux)
		params.AddIntSlice("matched_id", opts.MatchedID)
		params.AddStringSlice("matched_symbol", opts.MatchedSymbol)
		if opts.Category != nil {
			params.Add("category", string(*opts.Category))
		}
		if opts.FeeType != nil {
			params.Add("fee_type", string(*opts.FeeType))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
	}

	return get[map[string][]MarketPair](c, ctx, "/v2/cryptocurrency/market-pairs/latest", &RequestOptions[map[string][]MarketPair]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyOHLCVOptions struct {
	ID          []int
	Symbol      []string
	Convert     []string
	ConvertID   []int
	SkipInvalid *bool
}

func (c *Client) GetCryptocurrencyOHLCVLatest(ctx context.Context, opts *CryptocurrencyOHLCVOptions) (*APIResponse[map[string]OHLCV], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("symbol", opts.Symbol)
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
		params.AddBool("skip_invalid", opts.SkipInvalid)
	}

	return get[map[string]OHLCV](c, ctx, "/v2/cryptocurrency/ohlcv/latest", &RequestOptions[map[string]OHLCV]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyOHLCVHistoricalOptions struct {
	ID         []int
	Slug       []string
	Symbol     []string
	TimePeriod *string
	TimeStart  *string
	TimeEnd    *string
	Count      *int
	Interval   *Interval
	Convert    []string
	ConvertID  []int
}

func (c *Client) GetCryptocurrencyOHLCVHistorical(ctx context.Context, opts *CryptocurrencyOHLCVHistoricalOptions) (*APIResponse[map[string][]OHLCV], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.AddStringSlice("symbol", opts.Symbol)
		params.Add("time_period", *opts.TimePeriod)
		params.Add("time_start", *opts.TimeStart)
		params.Add("time_end", *opts.TimeEnd)
		params.AddInt("count", opts.Count)
		if opts.Interval != nil {
			params.Add("interval", string(*opts.Interval))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
	}

	return get[map[string][]OHLCV](c, ctx, "/v2/cryptocurrency/ohlcv/historical", &RequestOptions[map[string][]OHLCV]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyPricePerformanceStatsOptions struct {
	ID         []int
	Slug       []string
	Symbol     []string
	TimePeriod *TimePeriod
	Convert    []string
	ConvertID  []int
}

func (c *Client) GetCryptocurrencyPricePerformanceStats(ctx context.Context, opts *CryptocurrencyPricePerformanceStatsOptions) (*APIResponse[map[string]PricePerformanceStats], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.AddStringSlice("symbol", opts.Symbol)
		if opts.TimePeriod != nil {
			params.Add("time_period", string(*opts.TimePeriod))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
	}

	return get[map[string]PricePerformanceStats](c, ctx, "/v2/cryptocurrency/price-performance-stats/latest", &RequestOptions[map[string]PricePerformanceStats]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyCategoriesOptions struct {
	Start  *int
	Limit  *int
	ID     []int
	Slug   []string
	Symbol []string
}

func (c *Client) GetCryptocurrencyCategories(ctx context.Context, opts *CryptocurrencyCategoriesOptions) (*APIResponse[[]Category], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.AddStringSlice("symbol", opts.Symbol)
	}

	return get[[]Category](c, ctx, "/v1/cryptocurrency/categories", &RequestOptions[[]Category]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyCategoryOptions struct {
	ID      string
	Start   *int
	Limit   *int
	Convert []string
}

func (c *Client) GetCryptocurrencyCategory(ctx context.Context, opts *CryptocurrencyCategoryOptions) (*APIResponse[CategoryDetail], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.Add("id", opts.ID)
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.AddStringSlice("convert", opts.Convert)
	}

	return get[CategoryDetail](c, ctx, "/v1/cryptocurrency/category", &RequestOptions[CategoryDetail]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyAirdropsOptions struct {
	Start  *int
	Limit  *int
	Status *AirdropStatus
	ID     *int
	Slug   *string
	Symbol *string
}

func (c *Client) GetCryptocurrencyAirdrops(ctx context.Context, opts *CryptocurrencyAirdropsOptions) (*APIResponse[[]Airdrop], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.Status != nil {
			params.Add("status", string(*opts.Status))
		}
		params.AddInt("id", opts.ID)
		params.Add("slug", *opts.Slug)
		params.Add("symbol", *opts.Symbol)
	}

	return get[[]Airdrop](c, ctx, "/v1/cryptocurrency/airdrops", &RequestOptions[[]Airdrop]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetCryptocurrencyAirdrop(ctx context.Context, id string) (*APIResponse[Airdrop], error) {
	params := NewParamBuilder().Add("id", id)

	return get[Airdrop](c, ctx, "/v1/cryptocurrency/airdrop", &RequestOptions[Airdrop]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyTrendingOptions struct {
	Start      *int
	Limit      *int
	TimePeriod *TimePeriod
	Convert    []string
}

func (c *Client) GetCryptocurrencyTrendingLatest(ctx context.Context, opts *CryptocurrencyTrendingOptions) (*APIResponse[[]Trending], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.TimePeriod != nil {
			params.Add("time_period", string(*opts.TimePeriod))
		}
		params.AddStringSlice("convert", opts.Convert)
	}

	return get[[]Trending](c, ctx, "/v1/cryptocurrency/trending/latest", &RequestOptions[[]Trending]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetCryptocurrencyTrendingMostVisited(ctx context.Context, opts *CryptocurrencyTrendingOptions) (*APIResponse[[]Trending], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.TimePeriod != nil {
			params.Add("time_period", string(*opts.TimePeriod))
		}
		params.AddStringSlice("convert", opts.Convert)
	}

	return get[[]Trending](c, ctx, "/v1/cryptocurrency/trending/most-visited", &RequestOptions[[]Trending]{
		QueryParams: params.Build(),
	})
}

type CryptocurrencyGainersLosersOptions struct {
	Start      *int
	Limit      *int
	TimePeriod *TimePeriod
	Convert    []string
	Sort       *string
	SortDir    *SortDirection
}

func (c *Client) GetCryptocurrencyTrendingGainersLosers(ctx context.Context, opts *CryptocurrencyGainersLosersOptions) (*APIResponse[[]Trending], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.TimePeriod != nil {
			params.Add("time_period", string(*opts.TimePeriod))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.Add("sort", *opts.Sort)
		if opts.SortDir != nil {
			params.Add("sort_dir", string(*opts.SortDir))
		}
	}

	return get[[]Trending](c, ctx, "/v1/cryptocurrency/trending/gainers-losers", &RequestOptions[[]Trending]{
		QueryParams: params.Build(),
	})
}
