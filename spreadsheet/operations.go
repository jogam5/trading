package spreadsheet

import (
	//"bufio"
	"trading/models"
	//"encoding/csv"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/order"
	//"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"gopkg.in/Iwark/spreadsheet.v2"
	//"io"
	//"io/ioutil"
	"log"
	//"os"
	//"path/filepath"
	"sort"
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
func GetCandles(bitfinex *rest.Client, coin string, sheet *spreadsheet.Sheet) []models.Candle {
	candlesHist, _ := bitfinex.Candles.History("tETHUSD", "1h")
	//log.Println(candles.Snapshot[0])
	//log.Println(timestampToTime(candles.Snapshot[0].MTS))
	id := 1
	candles := []models.Candle{}
	for _, v := range candlesHist.Snapshot {
		//log.Println(v)
		var c models.Candle
		c = models.Candle{
			Id:        id,
			Timestamp: Int64ToS(v.MTS),
			Time:      timestampToTime(v.MTS),
			Open:      FToS(v.Open),
			Close:     FToS(v.Close),
		}
		candles = append(candles, c)
		id = id + 1
	}
	sort.Sort(sort.Reverse(models.ById(candles)))
	return candles
}

/*
==
Receive candles and write them in spreadsheet
==
*/
func WriteCandles(candles []models.Candle, sheet *spreadsheet.Sheet) {
	beginRow := int(ReturnLastCell(0, sheet).Row) + 1
	for _, candle := range candles {
		toFind := candle.Timestamp
		found := false
		for _, row := range sheet.Rows {
			for _, cell := range row {
				if cell.Value == toFind {
					found = true
					log.Println("Candle already stored")
				}
			}
		}

		if found == false {
			/* Write new candle */
			sheet.Update(beginRow, 0, candle.Timestamp)
			sheet.Update(beginRow, 1, candle.Time)
			sheet.Update(beginRow, 2, candle.Open)
			sheet.Update(beginRow, 3, candle.Close)
			beginRow += 1
		}
	}
	sheet.Synchronize()
}

/*
==
Place buy and sell orders in Bitfinex depending on the Moving
Average 20 Day strategy
To do:
1. Get latest Bid Price (or Ask)
2. Refactor code for submit order
3. Send notification (SMS, Email, other) for V 3.0
==
*/

func MovingAverage(bitfinex *rest.Client, bfxPub *rest.Client, positions []models.Position) {
	//1. Get data from spreadsheet
	data := positions[len(positions)-1]
	r, _ := bfxPub.Tickers.Get("tETHUSD")
	price := r.LastPrice
	log.Println(price)

	if !data.Rebalance { // ! for Development
		log.Println("---> MOVING AVERAGE: Rebalance position")
		// 2. Rebalance current position
		if data.ETH > data.USD {
			log.Println("---> Sell ETH, Buy USD")
			response, err := bitfinex.Orders.SubmitOrder(&order.NewRequest{
				Symbol: "tETHUSD",
				CID:    time.Now().Unix() / 1000,
				Amount: -SToF(data.ETH),//-0.02, // Change
				Type:   "EXCHANGE LIMIT",
				Price:  500, // Change
			})
			if err != nil {
				panic(err)
			} else {
				log.Println(response)
			}
			orders := response.NotifyInfo.(*order.Snapshot)
			//log.Println(response.NotifyInfo.(*order.Snapshot))
			var orderID int64
			for _, o := range orders.Snapshot {
				orderID = o.ID
			}
			log.Println(orderID)

		} else {
			log.Println("---> Buy ETH, Sell USD")
			response, err := bitfinex.Orders.SubmitOrder(&order.NewRequest{
				Symbol: "tETHUSD",
				CID:    time.Now().Unix() / 1000,
				Amount: SToF(data.USD)/price,//0.02, // Change
				Type:   "EXCHANGE LIMIT",
				Price:  100, // Change
			})
			if err != nil {
				panic(err)
			} else {
				log.Println(response)
			}
			orders := response.NotifyInfo.(*order.Snapshot)
			//log.Println(response.NotifyInfo.(*order.Snapshot))
			var orderID int64
			for _, o := range orders.Snapshot {
				orderID = o.ID
			}
			log.Println(orderID)
		}
		// 4. Notify V 3.0
	} else {
		log.Println("---> MOVING AVERAGE: Hodl position")
		// 5. Notify V 3.0
	}
	
}

/*func orderStatus(bitfinex *rest.Client, orderID int) {
	if orderID != "" {
		bfxPriv.Orders.GetByOrderId(52003067598)
	} else {
		log.Println("Order does not exist")
	}
}*/
