package coinmarketcap

import "context"

type ExchangeMapOptions struct {
	ListingStatus *ListingStatus
	Slug          []string
	Start         *int
	Limit         *int
	Sort          *ExchangeSort
	Aux           []string
	CryptoID      []int
}

func (c *Client) GetExchangeMap(ctx context.Context, opts *ExchangeMapOptions) (*APIResponse[[]ExchangeMap], error) {
	params := NewParamBuilder()

	if opts != nil {
		if opts.ListingStatus != nil {
			params.Add("listing_status", string(*opts.ListingStatus))
		}
		params.AddStringSlice("slug", opts.Slug)
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.Sort != nil {
			params.Add("sort", string(*opts.Sort))
		}
		params.AddStringSlice("aux", opts.Aux)
		params.AddIntSlice("crypto_id", opts.CryptoID)
	}

	return get[[]ExchangeMap](c, ctx, "/v1/exchange/map", &RequestOptions[[]ExchangeMap]{
		QueryParams: params.Build(),
	})
}

type ExchangeInfoOptions struct {
	ID   []int
	Slug []string
	Aux  []string
}

func (c *Client) GetExchangeInfo(ctx context.Context, opts *ExchangeInfoOptions) (*APIResponse[map[string]ExchangeInfo], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[map[string]ExchangeInfo](c, ctx, "/v1/exchange/info", &RequestOptions[map[string]ExchangeInfo]{
		QueryParams: params.Build(),
	})
}

type ExchangeListingsOptions struct {
	Start      *int
	Limit      *int
	Sort       *ExchangeSort
	SortDir    *SortDirection
	MarketType *MarketType
	Category   *ExchangeCategory
	Aux        []string
	Convert    []string
}

func (c *Client) GetExchangeListingsLatest(ctx context.Context, opts *ExchangeListingsOptions) (*APIResponse[[]ExchangeListing], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.Sort != nil {
			params.Add("sort", string(*opts.Sort))
		}
		if opts.SortDir != nil {
			params.Add("sort_dir", string(*opts.SortDir))
		}
		if opts.MarketType != nil {
			params.Add("market_type", string(*opts.MarketType))
		}
		if opts.Category != nil {
			params.Add("category", string(*opts.Category))
		}
		params.AddStringSlice("aux", opts.Aux)
		params.AddStringSlice("convert", opts.Convert)
	}

	return get[[]ExchangeListing](c, ctx, "/v1/exchange/listings/latest", &RequestOptions[[]ExchangeListing]{
		QueryParams: params.Build(),
	})
}

type ExchangeQuotesOptions struct {
	ID      []int
	Slug    []string
	Convert []string
	Aux     []string
}

func (c *Client) GetExchangeQuotesLatest(ctx context.Context, opts *ExchangeQuotesOptions) (*APIResponse[map[string]ExchangeQuote], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.AddStringSlice("convert", opts.Convert)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[map[string]ExchangeQuote](c, ctx, "/v1/exchange/quotes/latest", &RequestOptions[map[string]ExchangeQuote]{
		QueryParams: params.Build(),
	})
}

type ExchangeQuotesHistoricalOptions struct {
	ID        []int
	Slug      []string
	TimeStart *string
	TimeEnd   *string
	Count     *int
	Interval  *Interval
	Convert   []string
	Aux       []string
}

func (c *Client) GetExchangeQuotesHistorical(ctx context.Context, opts *ExchangeQuotesHistoricalOptions) (*APIResponse[map[string][]HistoricalQuote], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("slug", opts.Slug)
		params.Add("time_start", *opts.TimeStart)
		params.Add("time_end", *opts.TimeEnd)
		params.AddInt("count", opts.Count)
		if opts.Interval != nil {
			params.Add("interval", string(*opts.Interval))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[map[string][]HistoricalQuote](c, ctx, "/v1/exchange/quotes/historical", &RequestOptions[map[string][]HistoricalQuote]{
		QueryParams: params.Build(),
	})
}

type ExchangeMarketPairsOptions struct {
	ID            *int
	Slug          *string
	Start         *int
	Limit         *int
	Aux           []string
	MatchedID     []int
	MatchedSymbol []string
	Category      *PairCategory
	FeeType       *FeeType
	Convert       []string
}

func (c *Client) GetExchangeMarketPairsLatest(ctx context.Context, opts *ExchangeMarketPairsOptions) (*APIResponse[[]MarketPair], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("id", opts.ID)
		params.Add("slug", *opts.Slug)
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
	}

	return get[[]MarketPair](c, ctx, "/v1/exchange/market-pairs/latest", &RequestOptions[[]MarketPair]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetExchangeAssets(ctx context.Context, id int) (*APIResponse[map[string]interface{}], error) {
	params := NewParamBuilder().AddInt("id", &id)

	return get[map[string]interface{}](c, ctx, "/v1/exchange/assets", &RequestOptions[map[string]interface{}]{
		QueryParams: params.Build(),
	})
}
