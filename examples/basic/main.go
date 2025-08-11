package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Davincible/go-coinmarketcap"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("CMC_API_KEY")
	if apiKey == "" {
		log.Fatal("CMC_API_KEY environment variable is required")
	}

	// Create client
	client := coinmarketcap.NewClient(
		coinmarketcap.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	// Example 1: Get latest cryptocurrency listings
	fmt.Println("=== Latest Cryptocurrency Listings ===")
	listings, err := client.GetCryptocurrencyListingsLatest(ctx, &coinmarketcap.CryptocurrencyListingsOptions{
		Start:   coinmarketcap.Int(1),
		Limit:   coinmarketcap.Int(10),
		Convert: []string{"USD"},
	})
	if err != nil {
		log.Printf("Error getting listings: %v", err)
		return
	}

	for _, crypto := range listings.Data {
		price := "N/A"
		if quote, exists := crypto.Quote["USD"]; exists && quote.Price != nil {
			price = fmt.Sprintf("$%.2f", *quote.Price)
		}
		fmt.Printf("%s (%s): %s\n", crypto.Name, crypto.Symbol, price)
	}

	// Example 2: Get specific cryptocurrency quotes
	fmt.Println("\n=== Bitcoin and Ethereum Quotes ===")
	quotes, err := client.GetCryptocurrencyQuotesLatest(ctx, &coinmarketcap.CryptocurrencyQuotesOptions{
		Symbol:  []string{"BTC", "ETH"},
		Convert: []string{"USD", "EUR"},
	})
	if err != nil {
		log.Printf("Error getting quotes: %v", err)
		return
	}

	for _, quote := range quotes.Data {
		fmt.Printf("%s (%s):\n", quote.Name, quote.Symbol)
		for currency, q := range quote.Quote {
			if q.Price != nil {
				fmt.Printf("  %s: $%.2f\n", currency, *q.Price)
			}
		}
	}

	// Example 3: Get global market metrics
	fmt.Println("\n=== Global Market Metrics ===")
	globals, err := client.GetGlobalMetricsLatest(ctx, &coinmarketcap.GlobalMetricsOptions{
		Convert: []string{"USD"},
	})
	if err != nil {
		log.Printf("Error getting global metrics: %v", err)
		return
	}

	global := globals.Data
	if quote, exists := global.Quote["USD"]; exists {
		if quote.MarketCap != nil {
			fmt.Printf("Total Market Cap: $%.0f\n", *quote.MarketCap)
		}
		if quote.Volume24h != nil {
			fmt.Printf("24h Volume: $%.0f\n", *quote.Volume24h)
		}
	}

	if global.BtcDominance != nil {
		fmt.Printf("Bitcoin Dominance: %.2f%%\n", *global.BtcDominance)
	}

	// Example 4: Get trending cryptocurrencies
	fmt.Println("\n=== Trending Cryptocurrencies ===")
	trending, err := client.GetCryptocurrencyTrendingLatest(ctx, &coinmarketcap.CryptocurrencyTrendingOptions{
		Limit:   coinmarketcap.Int(5),
		Convert: []string{"USD"},
	})
	if err != nil {
		log.Printf("Error getting trending: %v", err)
		return
	}

	for i, trend := range trending.Data {
		price := "N/A"
		if quote, exists := trend.Quote["USD"]; exists && quote.Price != nil {
			price = fmt.Sprintf("$%.2f", *quote.Price)
		}
		fmt.Printf("%d. %s (%s): %s\n", i+1, trend.Name, trend.Symbol, price)
	}

	fmt.Printf("\nAPI Credits Used: %d\n", listings.Status.CreditCount)
}
