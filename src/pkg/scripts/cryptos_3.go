package scripts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// --- API response structs (simplified) ---

type APIResponse struct {
	Data []Crypto3 `json:"data"`
}

type Crypto3 struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  Quote  `json:"quote"`
}

type Quote struct {
	USD USDQuote `json:"USD"`
}

type USDQuote struct {
	Price     float64 `json:"price"`
	MarketCap float64 `json:"market_cap"`
}

// --- Main ---

func Info3() {

	godotenv.Load(".env")

	apiKey := os.Getenv("CMC_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set CMC_API_KEY environment variable")
		return
	}

	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=5"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		panic(err)
	}

	// Print results
	for _, c := range apiResp.Data {
		fmt.Printf(
			"%s (%s)\n  Price: $%.2f\n  Market Cap: $%.2f\n\n",
			c.Name,
			c.Symbol,
			c.Quote.USD.Price,
			c.Quote.USD.MarketCap,
		)
	}
}
