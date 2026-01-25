// Obtengo precio de cryptos classificado

package scripts

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/Finnhub-Stock-API/finnhub-go/v2"

	"github.com/joho/godotenv"
)

type Crypto struct {
	Symbol string
	Price  float64
}

func Info2() {

	godotenv.Load(".env")

	apiKey := os.Getenv("FINNHUB_API_KEY")
	if apiKey == "" {
		log.Fatal("FINNHUB_API_KEY not set")
	}

	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	client := finnhub.NewAPIClient(cfg)

	cryptos := []string{
		"BINANCE:BTCUSDT",
		"BINANCE:ETHUSDT",
		"BINANCE:SOLUSDT",
		"BINANCE:ADAUSDT",
	}

	var results []Crypto

	fmt.Println("📊 Fetching current crypto prices... ")

	for _, symbol := range cryptos {
		quote, _, err := client.DefaultApi.Quote(context.Background()).
			Symbol(symbol).
			Execute()
		if err != nil {
			log.Println("quote error:", symbol, err)
			continue
		}

		results = append(results, Crypto{
			Symbol: symbol,
			Price:  float64(*quote.C),
		})
	}

	// ---- Sort by price (descending) ----
	sort.Slice(results, func(i, j int) bool {
		return results[i].Price > results[j].Price
	})

	// ---- Classify by price ----
	fmt.Println("💰 Crypto classification (by price):")
	for _, c := range results {
		category := classify(c.Price)
		fmt.Printf("%-18s  $%-10.2f  %s\n", c.Symbol, c.Price, category)
	}

	// ---- Historical candles (last 5 days) ----
	fmt.Println("\n📈 5-day historical candles: ")

	to := time.Now().Unix()
	from := time.Now().AddDate(0, 0, -10).Unix()

	for _, c := range results {
		fmt.Println("🔹", c.Symbol)

		candles, _, err := client.DefaultApi.CryptoCandles(context.Background()).
			Symbol(c.Symbol).
			Resolution("60").
			From(from).
			To(to).
			Execute()

		if err != nil || *candles.S != "ok" {
			//fmt.Println("Status:", candles.S) // "no_data", "error"
			fmt.Println("  No candle data")
			continue
		}

		r := *candles.O
		s := *candles.C

		for i, v := range *candles.T {

			t := time.Unix(int64(v), 0).UTC()
			fmt.Printf(
				"  %s  O: %.2f  C: %.2f\n",
				t.Format("2006-01-02"),
				r[i],
				s[i],
			)

		}
		fmt.Println()
	}
}

// ---- Price classification ----
func classify(price float64) string {
	switch {
	case price >= 10000:
		return "🟢 High-value crypto"
	case price >= 100:
		return "🟡 Mid-value crypto"
	default:
		return "🔵 Low-value crypto"
	}
}
