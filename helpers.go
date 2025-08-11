package coinmarketcap

// Utility functions for working with pointer types in API options

// String returns a pointer to the string value.
// This is useful for setting optional string parameters.
func String(s string) *string {
	return &s
}

// Int returns a pointer to the int value.
// This is useful for setting optional integer parameters.
func Int(i int) *int {
	return &i
}

// Float64 returns a pointer to the float64 value.
// This is useful for setting optional float parameters.
func Float64(f float64) *float64 {
	return &f
}

// Bool returns a pointer to the bool value.
// This is useful for setting optional boolean parameters.
func Bool(b bool) *bool {
	return &b
}

// ListingSort returns a pointer to the ListingSort value.
func ListingSortPtr(s ListingSort) *ListingSort {
	return &s
}

// SortDirection returns a pointer to the SortDirection value.
func SortDirectionPtr(s SortDirection) *SortDirection {
	return &s
}

// CryptocurrencyType returns a pointer to the CryptocurrencyType value.
func CryptocurrencyTypePtr(t CryptocurrencyType) *CryptocurrencyType {
	return &t
}

// ListingStatus returns a pointer to the ListingStatus value.
func ListingStatusPtr(s ListingStatus) *ListingStatus {
	return &s
}

// Interval returns a pointer to the Interval value.
func IntervalPtr(i Interval) *Interval {
	return &i
}

// TimePeriod returns a pointer to the TimePeriod value.
func TimePeriodPtr(t TimePeriod) *TimePeriod {
	return &t
}

// MarketType returns a pointer to the MarketType value.
func MarketTypePtr(m MarketType) *MarketType {
	return &m
}

// ExchangeCategory returns a pointer to the ExchangeCategory value.
func ExchangeCategoryPtr(e ExchangeCategory) *ExchangeCategory {
	return &e
}

// FeeType returns a pointer to the FeeType value.
func FeeTypePtr(f FeeType) *FeeType {
	return &f
}

// PairCategory returns a pointer to the PairCategory value.
func PairCategoryPtr(p PairCategory) *PairCategory {
	return &p
}

// AirdropStatus returns a pointer to the AirdropStatus value.
func AirdropStatusPtr(a AirdropStatus) *AirdropStatus {
	return &a
}

// ExchangeSort returns a pointer to the ExchangeSort value.
func ExchangeSortPtr(e ExchangeSort) *ExchangeSort {
	return &e
}
