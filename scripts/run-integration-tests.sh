#!/bin/bash

# Script to run integration tests for CoinMarketCap Go SDK
# Requires CMC_API_KEY environment variable to be set

set -e

echo "CoinMarketCap Go SDK - Integration Tests"
echo "========================================"

# Check if API key is set
if [ -z "$CMC_API_KEY" ]; then
    echo "❌ Error: CMC_API_KEY environment variable is not set"
    echo ""
    echo "To run integration tests, you need a CoinMarketCap API key."
    echo "Get one from: https://coinmarketcap.com/api/"
    echo ""
    echo "Then set it as an environment variable:"
    echo "export CMC_API_KEY=your-api-key-here"
    echo ""
    echo "Or run with the key inline:"
    echo "CMC_API_KEY=your-api-key-here ./scripts/run-integration-tests.sh"
    exit 1
fi

echo "✅ CMC_API_KEY is set"
echo ""

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ Error: go.mod not found. Please run this script from the project root directory."
    exit 1
fi

echo "🔍 Running integration tests..."
echo "⚠️  Note: These tests make real API calls and will consume API credits"
echo "⚠️  Tests are rate-limited to be conservative with API usage"
echo ""

# Run the integration tests
echo "Running: go test -tags=integration -v"
echo ""

go test -tags=integration -v

echo ""
echo "✅ Integration tests completed!"
echo ""
echo "💡 Tips:"
echo "  - Integration tests use real API calls and consume credits"
echo "  - Tests are rate-limited to avoid hitting API limits"
echo "  - Use sandbox mode for unlimited testing with mock data"
echo "  - Check your API usage at: https://coinmarketcap.com/api/account"