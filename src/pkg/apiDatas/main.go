// Obtengo precio de acciones para un determinado symbol ( ticker - empresa )

package apiDatas

/*
import (
	"context"
	"fmt"
	"log"
	"os"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	// Get your API key from ENV
	apiKey := os.Getenv("FINNHUB_API_KEY")
	if apiKey == "" {
		log.Fatal("FINNHUB_API_KEY not set")
	}

	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	finnhubClient := finnhub.NewAPIClient(cfg)

	// Call Quote endpoint
	symbol := "AAPL"
	quote, _, err := finnhubClient.DefaultApi.Quote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Fatal(err)
	}

	// Print results
	fmt.Printf("Symbol: ---------------- %s\n", symbol)
	fmt.Printf("Current price:---------- %.2f\n", *quote.C)
	fmt.Printf("Open price: ------------ %.2f\n", *quote.O)
	fmt.Printf("High price: ------------ %.2f\n", *quote.H)
	fmt.Printf("Low price: ------------- %.2f\n", *quote.L)
	fmt.Printf("Previous close: -------- %.2f\n", *quote.Pc)
}


*/
