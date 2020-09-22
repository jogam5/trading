package spreadsheet

import (
	//"bufio"
	"trading/models"
	//"encoding/csv"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"gopkg.in/Iwark/spreadsheet.v2"
	//"io"
	//"io/ioutil"
	"log"
	//"os"
	//"path/filepath"
	//"sort"
	"time"
)

/*
==
Check error generic function
==
*/
func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

/*
==
Convert Unix milliseconds to time.Time
==
*/
func timestampToTime(msTime int64) string {
	time := time.Unix(0, msTime*int64(1000000))
	return time.Format("01-02-2006 15:04:05")
}

/*
==
Fetch current price from Bitfinex and update table
==
*/
func GetPrice(bitfinex *rest.Client, coin string, sheet *spreadsheet.Sheet) []models.Candle {
	//r, e := bitfinex.Tickers.Get(coin)
	//checkError(e)
	candlesHist, _ := bitfinex.Candles.History("tETHUSD", "1h")
	//log.Println(candles.Snapshot[0])
	//log.Println(timestampToTime(candles.Snapshot[0].MTS))

	id := 1
	candles := []models.Candle{}
	for _, v := range candlesHist.Snapshot {
		log.Println(v)
		var c models.Candle
		c = models.Candle {
			Id: id,
			Timestamp: Int64ToS(v.MTS),
			Time: timestampToTime(v.MTS),
			Open: FToS(v.Open),
			Close: FToS(v.Close),
		}
		candles = append(candles, c)
		id = id + 1
	}
	log.Println(candles)
	return candles
}

		/*pair := "t" + v.Coin + v.Base
		log.Println(pair)

		r, e := bitfinex.Tickers.Get(pair)
		checkError(e)
		sheet.Update(v.Row, 12, ToS(r.LastPrice))
		i += 1
		if i%29 == 0 {
			time.Sleep(50 * time.Second)
		}
		*/
	//today := time.Now()
	//sheet.Update(0, 24, today.Format("01-02-2006 15:04:05"))
	//sheet.Synchronize()
