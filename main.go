package main

import (
	"crypto/rand"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var AlphaChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Numbers string = "1234567890"

func stockID(length int, chars string) string {
	words := len(chars)
	stockName := make([]byte, length)
	rand.Read(stockName) // generates len(b) random bytes
	for word := 0; word < length; word++ {
		stockName[word] = chars[int(stockName[word])%words]
	}
	return string(stockName)
}
func randNum(length int) string {
	var chars string = "123"
	words := len(chars)
	stockName := make([]byte, length)
	rand.Read(stockName) // generates len(b) random bytes
	for word := 0; word < length; word++ {
		stockName[word] = chars[int(stockName[word])%words]
	}
	return string(stockName)
}

type stock struct {
	ID     string  `json:"-"`
	Time   string  `json:"time"`
	Symbol string  `json:"symbol"`
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
	Volume int     `json:"volume"`
}

func orgTime() string {
	dTime := time.Now().Format(time.RFC3339)
	return string(dTime)
}

var stocks = []stock{
	{ID: "0", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "1", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "2", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "3", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "4", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "5", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "6", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "7", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "8", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
	{ID: "9", Time: orgTime(), Symbol: stockID(4, AlphaChars), Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 1300000},
}

func getUpdateById(id string) (*stock, error) {
	for i, t := range stocks {
		if t.ID == id {
			return &stocks[i], nil
		}
	}
	return nil, errors.New("stock not found")
}

func stockRoutine() {
	// id := data.Param("id")
	var newCloseValue float32
	randId := stockID(1, Numbers)
	stock, _ := getUpdateById(randId)
	// fmt.Println(stock)

	//change the values
	// - Update `time` to current
	stock.Time = time.Now().Format(time.RFC3339)

	// - Pick a random number between -10% and +10% of `close` price and add it to `close` price
	randStockValPercent := stockID(1, Numbers)
	intVar, _ := strconv.Atoi(randStockValPercent)
	randStockVal := (stock.Close / 100) * float32(intVar)
	negPos := stockID(1, "12")
	if negPos == "1" {
		newCloseValue = stock.Close + float32(randStockVal)
	} else {
		newCloseValue = stock.Close - float32(randStockVal)
	}

	// - If new `close` is higher than the `high`, update the `high`
	// - If new `close` is lower than the `low`, update the `low`
	if newCloseValue > stock.High {
		stock.High = newCloseValue
	} else if newCloseValue < stock.Low {
		stock.Low = newCloseValue
	}

	// - Pick a random number between 0 to 1000 and add it to the volume
	stNumber := []byte(randNum(1))
	NumByteToInt, _ := strconv.Atoi(string(stNumber))
	intVol := stockID(NumByteToInt, Numbers)
	randVolume, _ := strconv.Atoi(intVol)
	stock.Volume = stock.Volume + randVolume

	// Ticks will be published as newline delimited JSON objects.
}

func bgTask() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for _ = range ticker.C {
		stockRoutine()
	}
}

func getStocks(data *gin.Context) {
	data.IndentedJSON(http.StatusOK, stocks)
}

func main() {
	// finalJson, err := json.MarshalIndent(stocks, "", "\t")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", finalJson)
	go bgTask()
	router := gin.Default()
	router.GET("/", getStocks)
	router.Run("localhost:9000")

}
