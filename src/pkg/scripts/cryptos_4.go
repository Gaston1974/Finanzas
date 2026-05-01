package scripts

import "fmt"

type AssetClass struct {
	Name        string
	MarketValue float64 // USD
}

func Info4() {
	// --- Market values in USD (example / aggregated) ---
	// Crypto from CoinMarketCap (total market cap)
	cryptoMarketCap := 1.65e12 // $1.65T

	// Stocks (global equity market cap)
	equitiesMarketCap := 105e12 // $105T

	// Bonds (global outstanding debt)
	bondsMarketCap := 133e12 // $133T

	// Commodities (gold, oil, etc.)
	commoditiesMarketCap := 25e12 // $25T

	assets := []AssetClass{
		{"Crypto", cryptoMarketCap},
		{"Equities", equitiesMarketCap},
		{"Bonds", bondsMarketCap},
		{"Commodities", commoditiesMarketCap},
	}

	var total float64
	for _, a := range assets {
		total += a.MarketValue
	}

	fmt.Printf("%-15s %-20s %-10s\n", "ASSET CLASS", "MARKET VALUE (USD)", "DOMINANCE")
	fmt.Println("------------------------------------------------------------")

	for _, a := range assets {
		dominance := (a.MarketValue / total) * 100
		fmt.Printf(
			"%-15s $%-19.2e %-9.2f%%\n",
			a.Name,
			a.MarketValue,
			dominance,
		)
	}
}
