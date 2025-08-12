// Package main demonstrates correct usage of the CoinMarketCap API client
// This example shows how to safely use the library without crashes or errors.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	coinmarketcap "github.com/Davincible/go-coinmarketcap"
)

func main() {
	apiKey := os.Getenv("CMC_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set CMC_API_KEY environment variable")
	}

	client := coinmarketcap.NewClient(
		coinmarketcap.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	// Example 1: Get cryptocurrency map with proper error handling
	fmt.Println("=== Testing GetCryptocurrencyMap ===")
	mapOpts := &coinmarketcap.CryptocurrencyMapOptions{
		Start: coinmarketcap.Int(1),
		Limit: coinmarketcap.Int(20),
		// Note: Sort is optional and can be nil
		ListingStatus: coinmarketcap.ListingStatusPtr(coinmarketcap.StatusActive),
	}

	mapResponse, err := client.GetCryptocurrencyMap(ctx, mapOpts)
	if err != nil {
		log.Printf("GetCryptocurrencyMap error: %v", err)
	} else if mapResponse != nil && mapResponse.Data != nil {
		fmt.Printf("Found %d cryptocurrencies\n", len(mapResponse.Data))
		for i, crypto := range mapResponse.Data {
			if i < 3 { // Show only first 3
				fmt.Printf("  - %s (%s) - ID: %d\n", crypto.Name, crypto.Symbol, crypto.ID)
			}
		}
	}

	// Example 2: Get cryptocurrency quotes with proper error handling
	fmt.Println("\n=== Testing GetCryptocurrencyQuotesLatest ===")
	quotesOpts := &coinmarketcap.CryptocurrencyQuotesOptions{
		Symbol:  []string{"BTC", "ETH"},
		Convert: []string{"USD"},
		Aux:     []string{"cmc_rank"},
	}

	quotesResponse, err := client.GetCryptocurrencyQuotesLatest(ctx, quotesOpts)
	if err != nil {
		log.Printf("GetCryptocurrencyQuotesLatest error: %v", err)
	} else if quotesResponse != nil && quotesResponse.Data != nil {
		fmt.Printf("Retrieved quotes for %d symbols\n", len(quotesResponse.Data))
		for symbol, quotes := range quotesResponse.Data {
			// Get the primary quote (first one) using our helper function
			quote := coinmarketcap.GetPrimaryQuote(quotes)
			if quote != nil && quote.Quote != nil && quote.Quote["USD"] != nil && quote.Quote["USD"].Price != nil {
				fmt.Printf("  - %s: $%.2f (Rank: %d) [%d matches]\n", 
					symbol, 
					*quote.Quote["USD"].Price,
					func() int { if quote.CMCRank != nil { return *quote.CMCRank } else { return 0 } }(),
					len(quotes))
			}
		}
	}

	// Example 3: Comprehensive project data function
	fmt.Println("\n=== Testing GetComprehensiveProjectData ===")
	projectData, err := GetComprehensiveProjectData(client, ctx, "bitcoin")
	if err != nil {
		log.Printf("GetComprehensiveProjectData error: %v", err)
	} else if projectData != nil {
		fmt.Printf("Project: %s (%s)\n", projectData.Name, projectData.Symbol)
		if projectData.Price != nil {
			fmt.Printf("Price: $%.2f\n", *projectData.Price)
		}
		if projectData.Rank != nil {
			fmt.Printf("CMC Rank: %d\n", *projectData.Rank)
		}
		if projectData.Description != nil && *projectData.Description != "" {
			fmt.Printf("Description: %.100s...\n", *projectData.Description)
		}
	}
}

// ComprehensiveProjectData holds all available data for a cryptocurrency project
type ComprehensiveProjectData struct {
	ID          int
	Name        string
	Symbol      string
	Slug        string
	Price       *float64
	MarketCap   *float64
	Rank        *int
	Description *string
	Logo        string
	URLs        map[string][]string
	Tags        []string
	DateAdded   string
	Platform    *coinmarketcap.Platform
}

// GetComprehensiveProjectData searches for a cryptocurrency by name/symbol and returns comprehensive data
func GetComprehensiveProjectData(client *coinmarketcap.Client, ctx context.Context, nameOrSymbol string) (*ComprehensiveProjectData, error) {
	// Step 1: Try to find the cryptocurrency by symbol first
	upperSymbol := strings.ToUpper(nameOrSymbol)
	
	// Get basic info and current quotes
	quotesOpts := &coinmarketcap.CryptocurrencyQuotesOptions{
		Symbol:  []string{upperSymbol},
		Convert: []string{"USD"},
		Aux:     []string{"cmc_rank"},
	}
	
	quotesResp, quotesErr := client.GetCryptocurrencyQuotesLatest(ctx, quotesOpts)
	var cryptoID int
	var basicData *coinmarketcap.CryptocurrencyQuote
	
	if quotesErr == nil && quotesResp != nil && quotesResp.Data != nil {
		for _, quotes := range quotesResp.Data {
			// Get the primary quote from the array
			quote := coinmarketcap.GetPrimaryQuote(quotes)
			if quote != nil && strings.EqualFold(quote.Symbol, upperSymbol) {
				cryptoID = quote.ID
				basicData = quote
				break
			}
		}
	}
	
	// If not found by symbol, try searching by name using map
	if cryptoID == 0 {
		mapOpts := &coinmarketcap.CryptocurrencyMapOptions{
			Start: coinmarketcap.Int(1),
			Limit: coinmarketcap.Int(5000), // Get more results for name search
		}
		
		mapResp, mapErr := client.GetCryptocurrencyMap(ctx, mapOpts)
		if mapErr != nil {
			return nil, fmt.Errorf("failed to search by name: %w", mapErr)
		}
		
		if mapResp != nil && mapResp.Data != nil {
			lowerNameOrSymbol := strings.ToLower(nameOrSymbol)
			for _, crypto := range mapResp.Data {
				if strings.EqualFold(crypto.Name, nameOrSymbol) || 
				   strings.EqualFold(crypto.Symbol, nameOrSymbol) ||
				   strings.Contains(strings.ToLower(crypto.Name), lowerNameOrSymbol) {
					cryptoID = crypto.ID
					
					// Get quotes for this found crypto
					quotesOpts := &coinmarketcap.CryptocurrencyQuotesOptions{
						ID:      []int{cryptoID},
						Convert: []string{"USD"},
						Aux:     []string{"cmc_rank"},
					}
					
					quotesResp, quotesErr = client.GetCryptocurrencyQuotesLatest(ctx, quotesOpts)
					if quotesErr == nil && quotesResp != nil && quotesResp.Data != nil {
						for _, quotes := range quotesResp.Data {
							// Get the primary quote from the array
							quote := coinmarketcap.GetPrimaryQuote(quotes)
							if quote != nil {
								basicData = quote
								break
							}
						}
					}
					break
				}
			}
		}
	}
	
	if cryptoID == 0 {
		return nil, fmt.Errorf("cryptocurrency not found: %s", nameOrSymbol)
	}
	
	// Step 2: Get detailed info
	infoOpts := &coinmarketcap.CryptocurrencyInfoOptions{
		ID: []int{cryptoID},
	}
	
	infoResp, infoErr := client.GetCryptocurrencyInfo(ctx, infoOpts)
	var detailedInfo *coinmarketcap.CryptocurrencyInfo
	
	if infoErr == nil && infoResp != nil && infoResp.Data != nil {
		for _, info := range infoResp.Data {
			detailedInfo = &info
			break
		}
	}
	
	// Step 3: Combine all data
	result := &ComprehensiveProjectData{
		ID: cryptoID,
	}
	
	if basicData != nil {
		result.Name = basicData.Name
		result.Symbol = basicData.Symbol
		result.Slug = basicData.Slug
		result.Rank = basicData.CMCRank
		result.Platform = basicData.Platform
		
		if basicData.DateAdded.IsZero() == false {
			result.DateAdded = basicData.DateAdded.Format("2006-01-02")
		}
		
		if basicData.Quote != nil && basicData.Quote["USD"] != nil {
			result.Price = basicData.Quote["USD"].Price
			result.MarketCap = basicData.Quote["USD"].MarketCap
		}
		
		result.Tags = basicData.Tags
	}
	
	if detailedInfo != nil {
		result.Description = detailedInfo.Description
		result.Logo = detailedInfo.Logo
		result.URLs = detailedInfo.URLs
		
		// Override with more detailed info if available
		if result.Name == "" {
			result.Name = detailedInfo.Name
		}
		if result.Symbol == "" {
			result.Symbol = detailedInfo.Symbol
		}
		if result.Slug == "" {
			result.Slug = detailedInfo.Slug
		}
		if len(result.Tags) == 0 {
			result.Tags = detailedInfo.Tags
		}
		if result.Platform == nil {
			result.Platform = detailedInfo.Platform
		}
		if result.DateAdded == "" && detailedInfo.DateAdded.IsZero() == false {
			result.DateAdded = detailedInfo.DateAdded.Format("2006-01-02")
		}
	}
	
	return result, nil
}