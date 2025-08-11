package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Davincible/go-coinmarketcap"
	"golang.org/x/time/rate"
)

func main() {
	apiKey := os.Getenv("CMC_API_KEY")
	if apiKey == "" {
		log.Fatal("CMC_API_KEY environment variable is required")
	}

	// Create client with custom configuration
	client := coinmarketcap.NewClient(
		coinmarketcap.WithAPIKey(apiKey),
		coinmarketcap.WithRateLimit(rate.Limit(2.0)), // 2 requests per second
		coinmarketcap.WithHTTPClient(&http.Client{
			Timeout: 30 * time.Second,
		}),
		coinmarketcap.WithUserAgent("MyApp/1.0"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Example 1: Advanced cryptocurrency listings with filters
	fmt.Println("=== Advanced Cryptocurrency Listings ===")
	listings, err := client.GetCryptocurrencyListingsLatest(ctx, &coinmarketcap.CryptocurrencyListingsOptions{
		Start:        coinmarketcap.Int(1),
		Limit:        coinmarketcap.Int(20),
		Convert:      []string{"USD", "BTC"},
		Sort:         coinmarketcap.ListingSortPtr(coinmarketcap.SortMarketCap),
		SortDir:      coinmarketcap.SortDirectionPtr(coinmarketcap.SortDesc),
		MarketCapMin: coinmarketcap.Float64(1000000), // Min $1M market cap
		PriceMin:     coinmarketcap.Float64(0.01),    // Min $0.01 price
	})
	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("Found %d cryptocurrencies matching criteria:\n", len(listings.Data))
	for i, crypto := range listings.Data[:5] { // Show top 5
		usdQuote := crypto.Quote["USD"]
		btcQuote := crypto.Quote["BTC"]

		fmt.Printf("%d. %s (%s)\n", i+1, crypto.Name, crypto.Symbol)
		if usdQuote != nil && usdQuote.Price != nil {
			fmt.Printf("   USD Price: $%.4f\n", *usdQuote.Price)
			if usdQuote.MarketCap != nil {
				fmt.Printf("   Market Cap: $%.0f\n", *usdQuote.MarketCap)
			}
		}
		if btcQuote != nil && btcQuote.Price != nil {
			fmt.Printf("   BTC Price: ₿%.8f\n", *btcQuote.Price)
		}
		fmt.Println()
	}

	// Example 2: Historical data analysis
	fmt.Println("=== Historical Data Analysis ===")
	startTime := time.Now().AddDate(0, 0, -7).Format(time.RFC3339) // 7 days ago
	endTime := time.Now().Format(time.RFC3339)

	historical, err := client.GetCryptocurrencyQuotesHistorical(ctx, &coinmarketcap.CryptocurrencyQuotesHistoricalOptions{
		Symbol:    []string{"BTC"},
		TimeStart: &startTime,
		TimeEnd:   &endTime,
		Interval:  coinmarketcap.IntervalPtr(coinmarketcap.IntervalDaily),
		Convert:   []string{"USD"},
	})
	if err != nil {
		handleError(err)
		return
	}

	if btcData, exists := historical.Data["BTC"]; exists {
		fmt.Printf("Bitcoin price history (last %d days):\n", len(btcData))
		for _, quote := range btcData {
			if usdQuote, exists := quote.Quote["USD"]; exists && usdQuote.Price != nil {
				fmt.Printf("  %s: $%.2f\n", quote.Timestamp.Format("2006-01-02"), *usdQuote.Price)
			}
		}
	}

	// Example 3: Exchange market analysis
	fmt.Println("\n=== Exchange Market Analysis ===")
	exchanges, err := client.GetExchangeListingsLatest(ctx, &coinmarketcap.ExchangeListingsOptions{
		Start:   coinmarketcap.Int(1),
		Limit:   coinmarketcap.Int(10),
		Sort:    coinmarketcap.ExchangeSortPtr(coinmarketcap.ExchangeSortVolume24h),
		SortDir: coinmarketcap.SortDirectionPtr(coinmarketcap.SortDesc),
		Convert: []string{"USD"},
	})
	if err != nil {
		handleError(err)
		return
	}

	fmt.Println("Top exchanges by 24h volume:")
	for i, exchange := range exchanges.Data {
		fmt.Printf("%d. %s\n", i+1, exchange.Name)
		if exchange.Volume24hReported != nil {
			fmt.Printf("   24h Volume: $%.0f\n", *exchange.Volume24hReported)
		}
		if exchange.NumMarketPairs != nil {
			fmt.Printf("   Market Pairs: %d\n", *exchange.NumMarketPairs)
		}
		fmt.Println()
	}

	// Example 4: Market pairs analysis for Bitcoin
	fmt.Println("=== Bitcoin Market Pairs ===")
	pairs, err := client.GetCryptocurrencyMarketPairsLatest(ctx, &coinmarketcap.CryptocurrencyMarketPairsOptions{
		Symbol:  coinmarketcap.String("BTC"),
		Start:   coinmarketcap.Int(1),
		Limit:   coinmarketcap.Int(5),
		Convert: []string{"USD"},
	})
	if err != nil {
		handleError(err)
		return
	}

	if btcPairs, exists := pairs.Data["BTC"]; exists {
		fmt.Printf("Top Bitcoin trading pairs:\n")
		for i, pair := range btcPairs {
			fmt.Printf("%d. %s on %s\n", i+1, pair.MarketPair, pair.ExchangeName)
			if usdQuote, exists := pair.Quote["USD"]; exists && usdQuote.Volume24h != nil {
				fmt.Printf("   24h Volume: $%.0f\n", *usdQuote.Volume24h)
			}
			fmt.Println()
		}
	}

	// Example 5: API key information
	fmt.Println("=== API Key Information ===")
	keyInfo, err := client.GetKeyInfo(ctx)
	if err != nil {
		handleError(err)
		return
	}

	info := keyInfo.Data
	fmt.Printf("Plan: %s\n", info.Plan.Name)
	fmt.Printf("Rate Limit: %d requests/minute\n", info.Plan.RateLimitMinute)
	fmt.Printf("Daily Credits Used: %d/%d\n",
		info.Usage.CurrentDay.CreditsUsed,
		info.Plan.CreditLimitDaily)
	fmt.Printf("Monthly Credits Used: %d/%d\n",
		info.Usage.CurrentMonth.CreditsUsed,
		info.Plan.CreditLimitMonthly)
}

func handleError(err error) {
	if apiErr, ok := err.(*coinmarketcap.APIError); ok {
		fmt.Printf("API Error: %s\n", apiErr.Error())

		if apiErr.IsRateLimit() {
			fmt.Println("Rate limit exceeded. Consider upgrading your plan or adding delays between requests.")
		} else if apiErr.IsAuthError() {
			fmt.Println("Authentication error. Check your API key.")
		} else if apiErr.IsPaymentRequired() {
			fmt.Println("Payment required. Please upgrade your plan.")
		}
	} else {
		log.Printf("Error: %v", err)
	}
}
