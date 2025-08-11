package coinmarketcap

import "context"

type GlobalMetricsOptions struct {
	Convert   []string
	ConvertID []int
}

func (c *Client) GetGlobalMetricsLatest(ctx context.Context, opts *GlobalMetricsOptions) (*APIResponse[GlobalMetrics], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
	}

	return get[GlobalMetrics](c, ctx, "/v1/global-metrics/quotes/latest", &RequestOptions[GlobalMetrics]{
		QueryParams: params.Build(),
	})
}

type GlobalMetricsHistoricalOptions struct {
	TimeStart *string
	TimeEnd   *string
	Count     *int
	Interval  *Interval
	Convert   []string
	Aux       []string
}

func (c *Client) GetGlobalMetricsHistorical(ctx context.Context, opts *GlobalMetricsHistoricalOptions) (*APIResponse[[]GlobalMetrics], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.Add("time_start", *opts.TimeStart)
		params.Add("time_end", *opts.TimeEnd)
		params.AddInt("count", opts.Count)
		if opts.Interval != nil {
			params.Add("interval", string(*opts.Interval))
		}
		params.AddStringSlice("convert", opts.Convert)
		params.AddStringSlice("aux", opts.Aux)
	}

	return get[[]GlobalMetrics](c, ctx, "/v1/global-metrics/quotes/historical", &RequestOptions[[]GlobalMetrics]{
		QueryParams: params.Build(),
	})
}

type FiatMapOptions struct {
	Start         *int
	Limit         *int
	Sort          *string
	IncludeMetals *bool
}

func (c *Client) GetFiatMap(ctx context.Context, opts *FiatMapOptions) (*APIResponse[[]FiatMap], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.Add("sort", *opts.Sort)
		params.AddBool("include_metals", opts.IncludeMetals)
	}

	return get[[]FiatMap](c, ctx, "/v1/fiat/map", &RequestOptions[[]FiatMap]{
		QueryParams: params.Build(),
	})
}

type PriceConversionOptions struct {
	Amount    float64
	ID        *int
	Symbol    *string
	Time      *string
	Convert   []string
	ConvertID []int
}

func (c *Client) GetPriceConversion(ctx context.Context, opts *PriceConversionOptions) (*APIResponse[PriceConversion], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddFloat("amount", &opts.Amount)
		params.AddInt("id", opts.ID)
		params.Add("symbol", *opts.Symbol)
		params.Add("time", *opts.Time)
		params.AddStringSlice("convert", opts.Convert)
		params.AddIntSlice("convert_id", opts.ConvertID)
	}

	return get[PriceConversion](c, ctx, "/v2/tools/price-conversion", &RequestOptions[PriceConversion]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetPostmanCollection(ctx context.Context) (*APIResponse[interface{}], error) {
	return get[interface{}](c, ctx, "/v1/tools/postman", &RequestOptions[interface{}]{})
}

type BlockchainStatsOptions struct {
	ID     []int
	Symbol []string
	Slug   []string
}

func (c *Client) GetBlockchainStatsLatest(ctx context.Context, opts *BlockchainStatsOptions) (*APIResponse[map[string]BlockchainStats], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddIntSlice("id", opts.ID)
		params.AddStringSlice("symbol", opts.Symbol)
		params.AddStringSlice("slug", opts.Slug)
	}

	return get[map[string]BlockchainStats](c, ctx, "/v1/blockchain/statistics/latest", &RequestOptions[map[string]BlockchainStats]{
		QueryParams: params.Build(),
	})
}

type ContentLatestOptions struct {
	Start            *int
	Limit            *int
	Category         *string
	CryptocurrencyID *int
	Language         *string
	Sort             *string
}

func (c *Client) GetContentLatest(ctx context.Context, opts *ContentLatestOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.Add("category", *opts.Category)
		params.AddInt("cryptocurrency_id", opts.CryptocurrencyID)
		params.Add("language", *opts.Language)
		params.Add("sort", *opts.Sort)
	}

	return get[[]interface{}](c, ctx, "/v1/content/latest", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}

type ContentPostsOptions struct {
	TimePeriod       *TimePeriod
	CryptocurrencyID *int
	Start            *int
	Limit            *int
	Sort             *string
}

func (c *Client) GetContentPostsTop(ctx context.Context, opts *ContentPostsOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		if opts.TimePeriod != nil {
			params.Add("time_period", string(*opts.TimePeriod))
		}
		params.AddInt("cryptocurrency_id", opts.CryptocurrencyID)
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.Add("sort", *opts.Sort)
	}

	return get[[]interface{}](c, ctx, "/v1/content/posts/top", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetContentPostsLatest(ctx context.Context, opts *ContentPostsOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("cryptocurrency_id", opts.CryptocurrencyID)
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		params.Add("sort", *opts.Sort)
	}

	return get[[]interface{}](c, ctx, "/v1/content/posts/latest", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}

type ContentCommentsOptions struct {
	PostID string
	Start  *int
	Limit  *int
}

func (c *Client) GetContentPostsComments(ctx context.Context, opts *ContentCommentsOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.Add("post_id", opts.PostID)
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
	}

	return get[[]interface{}](c, ctx, "/v1/content/posts/comments", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}

type CommunityTrendingOptions struct {
	Start      *int
	Limit      *int
	TimePeriod *TimePeriod
}

func (c *Client) GetCommunityTrendingTopic(ctx context.Context, opts *CommunityTrendingOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.TimePeriod != nil {
			params.Add("time_period", string(*opts.TimePeriod))
		}
	}

	return get[[]interface{}](c, ctx, "/v1/community/trending/topic", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetCommunityTrendingToken(ctx context.Context, opts *CommunityTrendingOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
		if opts.TimePeriod != nil {
			params.Add("time_period", string(*opts.TimePeriod))
		}
	}

	return get[[]interface{}](c, ctx, "/v1/community/trending/token", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetKeyInfo(ctx context.Context) (*APIResponse[KeyInfo], error) {
	return get[KeyInfo](c, ctx, "/v1/key/info", &RequestOptions[KeyInfo]{})
}

type IndexOptions struct {
	TimeStart *string
	TimeEnd   *string
	Count     *string
	Interval  *string
}

func (c *Client) GetIndexCMC100Latest(ctx context.Context) (*APIResponse[interface{}], error) {
	return get[interface{}](c, ctx, "/v3/index/cmc100-latest", &RequestOptions[interface{}]{})
}

func (c *Client) GetIndexCMC100Historical(ctx context.Context, opts *IndexOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.Add("time_start", *opts.TimeStart)
		params.Add("time_end", *opts.TimeEnd)
		params.Add("count", *opts.Count)
		params.Add("interval", *opts.Interval)
	}

	return get[[]interface{}](c, ctx, "/v3/index/cmc100-historical", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}

func (c *Client) GetFearAndGreedLatest(ctx context.Context) (*APIResponse[interface{}], error) {
	return get[interface{}](c, ctx, "/v3/fear-and-greed/latest", &RequestOptions[interface{}]{})
}

type FearAndGreedHistoricalOptions struct {
	Start *int
	Limit *int
}

func (c *Client) GetFearAndGreedHistorical(ctx context.Context, opts *FearAndGreedHistoricalOptions) (*APIResponse[[]interface{}], error) {
	params := NewParamBuilder()

	if opts != nil {
		params.AddInt("start", opts.Start)
		params.AddInt("limit", opts.Limit)
	}

	return get[[]interface{}](c, ctx, "/v3/fear-and-greed/historical", &RequestOptions[[]interface{}]{
		QueryParams: params.Build(),
	})
}
