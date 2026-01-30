package scripts

import (
	"encoding/json"
	"fmt"
	"hello/src/pkg/apiDatas"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type Kline struct {
	Symbol    string
	OpenTime  int64
	Open      string
	High      string
	Low       string
	Close     string
	Volume    string
	CloseTime int64
}

func Info1(path string, days, months, years string) (int, string) {

	// Cryptos to fetch
	cryptos := []string{"BTCUSDT", "ETHUSDT", "SOLUSDT", "ADAUSDT"}

	// Time range: last days

	daysF, _ := strconv.ParseInt(days, 10, 0)
	monthsF, _ := strconv.ParseInt(months, 10, 0)
	yearsF, _ := strconv.ParseInt(years, 10, 0)

	to := time.Now().UTC()
	from := to.AddDate(int(yearsF), int(monthsF), int(daysF))

	// Create a new Excel file
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf("Failed to add sheet: %s\n", err.Error())
		return 0, err.Error()
	}

	style := apiDatas.AddStyle(2)

	// Add row and cells

	roww := sheet.AddRow()

	cryp := roww.AddCell()
	cryp.SetStyle(style)
	cryp.Value = "CRYPTO"
	fecha := roww.AddCell()
	fecha.SetStyle(style)
	fecha.Value = "FECHA"
	open := roww.AddCell()
	open.SetStyle(style)
	open.Value = "APERTURA"
	high := roww.AddCell()
	high.SetStyle(style)
	high.Value = "MAXIMO"
	low := roww.AddCell()
	low.SetStyle(style)
	low.Value = "MINIMO"
	close := roww.AddCell()
	close.SetStyle(style)
	close.Value = "CIERRE"
	volume := roww.AddCell()
	volume.SetStyle(style)
	volume.Value = "VOLUMEN"

	var totalklines []Kline

	i := 0

	for _, symbol := range cryptos {
		//fmt.Println("🔹", symbol)

		klines, err := fetchBinanceKlines(symbol, "1d", from, to)
		totalklines = append(totalklines, klines...)

		if err != nil {
			log.Println("Error:", err)
			i = 1
			continue
		}

	}

	for _, v := range totalklines {

		style = apiDatas.AddStyle(1)
		style3 := apiDatas.AddStyle(3)
		roww = sheet.AddRow()
		maxColumn0Lenght, maxColumn1Lenght := 0.0, 0.0

		cryp := roww.AddCell()
		cryp.SetStyle(style)
		cryp.Value = v.Symbol
		if float64(len(cryp.Value)) > maxColumn0Lenght {
			maxColumn0Lenght = float64(len(cryp.Value))
			sheet.SetColWidth(0, 0, maxColumn0Lenght)
		}

		fecha := roww.AddCell()
		fecha.SetStyle(style)
		t := time.UnixMilli(v.OpenTime)
		fecha.Value = t.Format("2006-01-02")
		if float64(len(fecha.Value)) > maxColumn1Lenght {
			maxColumn1Lenght = float64(len(fecha.Value))
			sheet.SetColWidth(1, 1, maxColumn1Lenght)
		}

		open := roww.AddCell()
		open.SetStyle(style3)
		sheet.SetColWidth(3, 3, 12)
		num, _ := strconv.ParseFloat(v.Open, 64)
		open.SetFloat(num)
		open.NumFmt = "#,##0.00"

		high := roww.AddCell()
		high.SetStyle(style3)
		sheet.SetColWidth(4, 4, 10)
		num, _ = strconv.ParseFloat(v.High, 64)
		high.SetFloat(num)
		high.NumFmt = "#,##0.00"

		low := roww.AddCell()
		low.SetStyle(style3)
		sheet.SetColWidth(5, 5, 10)
		num, _ = strconv.ParseFloat(v.Low, 64)
		low.SetFloat(num)
		low.NumFmt = "#,##0.00"

		close := roww.AddCell()
		close.SetStyle(style3)
		sheet.SetColWidth(6, 6, 10)
		num, _ = strconv.ParseFloat(v.Close, 64)
		close.SetFloat(num)
		close.NumFmt = "#,##0.00"

		volume := roww.AddCell()
		volume.SetStyle(style3)
		sheet.SetColWidth(7, 7, 14)
		num, _ = strconv.ParseFloat(v.Volume, 64)
		volume.SetFloat(num)
		volume.NumFmt = "#,##0.00"

		// closeTime := roww.AddCell()
		// closeTime.SetStyle(style)
		// t = time.UnixMilli(v.CloseTime)
		// closeTime.Value = t.Format("2006-01-02")
		// if float64(len(closeTime.Value))*1.2 > maxColumn2Lenght {
		// 	maxColumn2Lenght = float64(len(closeTime.Value)) * 1.2
		// 	sheet.SetColWidth(8, 8, maxColumn2Lenght)
		// }

	}

	// Save the file
	err = file.Save(path + "/BINANCE.xlsx")
	if err != nil || i == 1 {
		fmt.Printf("Failed to save file: %s\n", err)
		return 1, "archivo generado"
	} else {
		fmt.Println("Excel file 'BINANCE.xlsx' created successfully.")
		return 0, "error en armado de excel"
	}

}

// Fetch Binance Klines
func fetchBinanceKlines(symbol, interval string, from, to time.Time) ([]Kline, error) {
	base := "https://api.binance.com/api/v3/klines"
	url := fmt.Sprintf("%s?symbol=%s&interval=%s&startTime=%d&endTime=%d",
		base,
		symbol,
		interval,
		from.UnixMilli(),
		to.UnixMilli(),
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var klines []Kline
	for _, r := range raw {
		klines = append(klines, Kline{
			Symbol:    symbol,
			OpenTime:  int64(r[0].(float64)),
			Open:      r[1].(string),
			High:      r[2].(string),
			Low:       r[3].(string),
			Close:     r[4].(string),
			Volume:    r[5].(string),
			CloseTime: int64(r[6].(float64)),
		})
	}

	return klines, nil
}
