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
1. Send notification (SMS, Email, other) for V 3.0
==
*/

func MovingAverage(sheet *spreadsheet.Sheet, bfxPriv *rest.Client, bfxPub *rest.Client, positions []models.Position) {
	//1. Get data from spreadsheet
	var orderID int64
	data := positions[len(positions)-1]
	r, _ := bfxPub.Tickers.Get("tETHUSD")

	// 2. Rebalance current position
	if !data.Rebalance { // ! for Development
		log.Println("---> MOVING AVERAGE: Rebalance position")
		if data.ETH > data.USD {
			log.Println("---> Sell ETH, Buy USD")
			price := r.Ask + 0.3
			log.Println(price)
			amount := -1 * (SToF(data.ETH))
			log.Println(amount)
			orderID = SubmitOrder(bfxPriv, price, amount)
		} else {
			log.Println("---> Buy ETH, Sell USD")
			price := r.Bid - 0.3
			log.Println("---> Price:", price)
			amount := SToF(data.USD) / price
			log.Println("---> Amount:", amount)
			orderID = SubmitOrder(bfxPriv, price, amount)
		}
		// 4. Notify V 3.0
	} else {
		log.Println("---> MOVING AVERAGE: Hodl position")
		// 5. Notify V 3.0
	}
	sheet.Update(data.Id, 12, Int64ToS(orderID))
	sheet.Synchronize()
}

/*
==
Receives the price and the amount in order to submit an order in
Bitfinex and returns the orderID, which can be used to check
the status of the submitted order.
==
*/

func SubmitOrder(bfxPriv *rest.Client, price float64, amount float64) int64 {
	response, err := bfxPriv.Orders.SubmitOrder(&order.NewRequest{
		Symbol: "tETHUSD",
		CID:    time.Now().Unix() / 1000,
		Amount: amount,
		Type:   "EXCHANGE LIMIT",
		Price:  price,
	})
	if err != nil {
		panic(err)
	} else {
		//log.Println(response)
	}
	orders := response.NotifyInfo.(*order.Snapshot)
	var orderID int64
	for _, o := range orders.Snapshot {
		orderID = o.ID
	}
	log.Println("---> OrderID:", orderID)
	return orderID
}

/*func orderStatus(bitfinex *rest.Client, orderID int) {
	if orderID != "" {
		bfxPriv.Orders.GetByOrderId(52003067598)
	} else {
		log.Println("Order does not exist")
	}
}*/
