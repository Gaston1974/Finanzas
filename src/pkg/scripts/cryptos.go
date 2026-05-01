// Obtengo precio de cryptos classificado

package scripts

import (
	"context"

	"log"
	"os"
	"sort"
	"strings"

	"fyne.io/fyne/v2/widget"
	"github.com/Finnhub-Stock-API/finnhub-go/v2"

	"github.com/joho/godotenv"
)

type Crypto struct {
	Symbol string
	Price  float64
}

type Classify struct {
	Name     Crypto
	Category string
}

func Info2(form *widget.Form, lbl *widget.Label) ([]Classify, int) {

	godotenv.Load(".env")

	ch := make(chan int)

	Worker(lbl, ch)

	apiKey := os.Getenv("FINNHUB_API_KEY")
	if apiKey == "" {
		log.Fatal("FINNHUB_API_KEY not set")
		return nil, 0
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

	// fmt.Println("📊 Fetching current crypto prices... ")

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

	var Str Classify
	var Info []Classify

	// ---- Classify by price ----
	// fmt.Println("💰 Crypto classification (by price):")
	for _, c := range results {

		category := classify(c.Price)
		// fmt.Printf("%-18s  $%-10.2f  %s\n", c.Symbol, c.Price, category)

		Str.Name.Price = c.Price
		_, aft, _ := strings.Cut(c.Symbol, ":")
		Str.Name.Symbol = aft
		Str.Category = category

		Info = append(Info, Str)

	}

	ch <- 1
	form.Refresh()
	close(ch) // tell the goroutine we're done

	return Info, 1

	// ---- Historical candles (last 5 days) ----
	// fmt.Println("\n📈 5-day historical candles: ")

	// to := time.Now().Unix()
	// from := time.Now().AddDate(0, 0, -10).Unix()

	// for _, c := range results {
	// 	fmt.Println("🔹", c.Symbol)

	// 	candles, _, err := client.DefaultApi.CryptoCandles(context.Background()).
	// 		Symbol(c.Symbol).
	// 		Resolution("60").
	// 		From(from).
	// 		To(to).
	// 		Execute()

	// 	if err != nil || *candles.S != "ok" {
	// 		//fmt.Println("Status:", candles.S) // "no_data", "error"
	// 		fmt.Println("  No candle data")
	// 		continue
	// 	}

	// 	r := *candles.O
	// 	s := *candles.C

	// 	for i, v := range *candles.T {

	// 		t := time.Unix(int64(v), 0).UTC()
	// 		fmt.Printf(
	// 			"  %s  O: %.2f  C: %.2f\n",
	// 			t.Format("2006-01-02"),
	// 			r[i],
	// 			s[i],
	// 		)

	// 	}
	// 	fmt.Println()
	// }
}

// ---- Price classification ----
func classify(price float64) string {
	switch {
	case price >= 10000:
		//return "🟢 High-value crypto"
		return "ALTO"
	case price >= 100:
		//return "🟡 Mid-value crypto"
		return "MEDIO"
	default:
		//return "🔵 Low-value crypto"
		return "BAJO"
	}
}

// func ProgressBar(a fyne.App, ch <-chan int) {

// 	w := a.NewWindow("")

// 	progress := widget.NewProgressBar()
// 	infinite := widget.NewProgressBarInfinite()

// 	go func() {
// 		for i := 0.0; i <= 1.0; i += 0.1 {
// 			time.Sleep(time.Millisecond * 250)
// 			progress.SetValue(i)

// 		}

// 		msg := <-ch

// 		if msg == 1 {
// 			w.Close()
// 		}

// 	}()

// 	w.SetContent(container.NewVBox(progress, infinite))
// 	w.Resize(fyne.NewSize(250, 50))
// 	w.Show()
// }

func Worker(lbl *widget.Label, ch <-chan int) {

	lbl.Text = "PROCESANDO.."

	go func() {

		msg := <-ch

		if msg == 1 {

			lbl.Text = ""

		}

	}()

}
