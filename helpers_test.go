package coinmarketcap

import "testing"

func TestHelperFunctions(t *testing.T) {
	// Test String helper
	str := String("test")
	if str == nil || *str != "test" {
		t.Errorf("expected String('test') to return pointer to 'test', got %v", str)
	}

	// Test Int helper
	intVal := Int(123)
	if intVal == nil || *intVal != 123 {
		t.Errorf("expected Int(123) to return pointer to 123, got %v", intVal)
	}

	// Test Float64 helper
	floatVal := Float64(12.34)
	if floatVal == nil || *floatVal != 12.34 {
		t.Errorf("expected Float64(12.34) to return pointer to 12.34, got %v", floatVal)
	}

	// Test Bool helper
	boolVal := Bool(true)
	if boolVal == nil || *boolVal != true {
		t.Errorf("expected Bool(true) to return pointer to true, got %v", boolVal)
	}

	boolVal2 := Bool(false)
	if boolVal2 == nil || *boolVal2 != false {
		t.Errorf("expected Bool(false) to return pointer to false, got %v", boolVal2)
	}

	// Test ListingSortPtr helper
	sortVal := ListingSortPtr(SortMarketCap)
	if sortVal == nil || *sortVal != SortMarketCap {
		t.Errorf("expected ListingSortPtr to return pointer to SortMarketCap, got %v", sortVal)
	}

	// Test SortDirectionPtr helper
	sortDir := SortDirectionPtr(SortAsc)
	if sortDir == nil || *sortDir != SortAsc {
		t.Errorf("expected SortDirectionPtr to return pointer to SortAsc, got %v", sortDir)
	}

	// Test CryptocurrencyTypePtr helper
	cryptoType := CryptocurrencyTypePtr(CryptoTypeAll)
	if cryptoType == nil || *cryptoType != CryptoTypeAll {
		t.Errorf("expected CryptocurrencyTypePtr to return pointer to CryptoTypeAll, got %v", cryptoType)
	}

	// Test ListingStatusPtr helper
	status := ListingStatusPtr(StatusActive)
	if status == nil || *status != StatusActive {
		t.Errorf("expected ListingStatusPtr to return pointer to StatusActive, got %v", status)
	}

	// Test IntervalPtr helper
	interval := IntervalPtr(Interval1h)
	if interval == nil || *interval != Interval1h {
		t.Errorf("expected IntervalPtr to return pointer to Interval1h, got %v", interval)
	}

	// Test TimePeriodPtr helper
	timePeriod := TimePeriodPtr(TimePeriod24h)
	if timePeriod == nil || *timePeriod != TimePeriod24h {
		t.Errorf("expected TimePeriodPtr to return pointer to TimePeriod24h, got %v", timePeriod)
	}

	// Test MarketTypePtr helper
	marketType := MarketTypePtr(MarketTypeAll)
	if marketType == nil || *marketType != MarketTypeAll {
		t.Errorf("expected MarketTypePtr to return pointer to MarketTypeAll, got %v", marketType)
	}

	// Test ExchangeCategoryPtr helper
	exchangeCategory := ExchangeCategoryPtr(ExchangeCategoryAll)
	if exchangeCategory == nil || *exchangeCategory != ExchangeCategoryAll {
		t.Errorf("expected ExchangeCategoryPtr to return pointer to ExchangeCategoryAll, got %v", exchangeCategory)
	}

	// Test FeeTypePtr helper
	feeType := FeeTypePtr(FeeTypeAll)
	if feeType == nil || *feeType != FeeTypeAll {
		t.Errorf("expected FeeTypePtr to return pointer to FeeTypeAll, got %v", feeType)
	}

	// Test PairCategoryPtr helper
	pairCategory := PairCategoryPtr(PairCategoryAll)
	if pairCategory == nil || *pairCategory != PairCategoryAll {
		t.Errorf("expected PairCategoryPtr to return pointer to PairCategoryAll, got %v", pairCategory)
	}

	// Test AirdropStatusPtr helper
	airdropStatus := AirdropStatusPtr(AirdropStatusOngoing)
	if airdropStatus == nil || *airdropStatus != AirdropStatusOngoing {
		t.Errorf("expected AirdropStatusPtr to return pointer to AirdropStatusOngoing, got %v", airdropStatus)
	}

	// Test ExchangeSortPtr helper
	exchangeSort := ExchangeSortPtr(ExchangeSortVolume24h)
	if exchangeSort == nil || *exchangeSort != ExchangeSortVolume24h {
		t.Errorf("expected ExchangeSortPtr to return pointer to ExchangeSortVolume24h, got %v", exchangeSort)
	}
}

func TestHelperPointerUniqueness(t *testing.T) {
	// Test that helper functions return unique pointers
	str1 := String("test")
	str2 := String("test")

	if str1 == str2 {
		t.Error("String() should return unique pointers for same values")
	}

	int1 := Int(123)
	int2 := Int(123)

	if int1 == int2 {
		t.Error("Int() should return unique pointers for same values")
	}

	float1 := Float64(12.34)
	float2 := Float64(12.34)

	if float1 == float2 {
		t.Error("Float64() should return unique pointers for same values")
	}

	bool1 := Bool(true)
	bool2 := Bool(true)

	if bool1 == bool2 {
		t.Error("Bool() should return unique pointers for same values")
	}
}

func TestHelperUsageInOptions(t *testing.T) {
	// Test that helpers work well in option structs
	opts := &CryptocurrencyListingsOptions{
		Start:              Int(1),
		Limit:              Int(10),
		PriceMin:           Float64(1.0),
		PriceMax:           Float64(1000.0),
		Sort:               ListingSortPtr(SortMarketCap),
		SortDir:            SortDirectionPtr(SortDesc),
		CryptocurrencyType: CryptocurrencyTypePtr(CryptoTypeCoins),
		Tag:                String("defi"),
	}

	if opts.Start == nil || *opts.Start != 1 {
		t.Error("Start option not set correctly")
	}
	if opts.Limit == nil || *opts.Limit != 10 {
		t.Error("Limit option not set correctly")
	}
	if opts.PriceMin == nil || *opts.PriceMin != 1.0 {
		t.Error("PriceMin option not set correctly")
	}
	if opts.Sort == nil || *opts.Sort != SortMarketCap {
		t.Error("Sort option not set correctly")
	}
	if opts.SortDir == nil || *opts.SortDir != SortDesc {
		t.Error("SortDir option not set correctly")
	}
	if opts.CryptocurrencyType == nil || *opts.CryptocurrencyType != CryptoTypeCoins {
		t.Error("CryptocurrencyType option not set correctly")
	}
	if opts.Tag == nil || *opts.Tag != "defi" {
		t.Error("Tag option not set correctly")
	}
}
