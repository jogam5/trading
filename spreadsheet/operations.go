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
	"strings"
	"time"
)

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
Fetch current price candles from Bitfinex and update table
==
*/
func GetCandles(bitfinex *rest.Client, coin string, sheet *spreadsheet.Sheet) []models.Candle {
	candlesHist, _ := bitfinex.Candles.History("tETHUSD", "1h")
	//log.Println(timestampToTime(candles.Snapshot[0].MTS))
	id := 1
	candles := []models.Candle{}
	for _, v := range candlesHist.Snapshot {
		var c models.Candle
		c = models.Candle{
			Id:        id,
			Timestamp: Int64ToString(v.MTS),
			Time:      timestampToTime(v.MTS),
			Open:      Float64ToString(v.Open),
			Close:     Float64ToString(v.Close),
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
					//log.Println("Candle already stored")
				}
			}
		}

		if found == false {
			/* Write new candle */
			log.Println("--> New candle:", candle)
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
==
*/
func MovingAverage(sheet *spreadsheet.Sheet, bfxPriv *rest.Client, bfxPub *rest.Client, positions []models.Position) int {
	/* 1. Get data from spreadsheet */
	var orderID int64
	var status string
	data := positions[len(positions)-1]
	r, _ := bfxPub.Tickers.Get("tETHUSD")

	if strings.Contains(data.Status, "EXECUTED") || strings.Contains(data.Status, "ACTIVE") {
		log.Println("--> Error: MA strategy was already applied on the last data.")
		return 1
	}

	/* 2. Rebalance current position */
	if data.Rebalance {
		log.Println("---> MOVING AVERAGE: Rebalance position")
		if data.ETH > data.USD {
			log.Println("---> Sell ETH, Buy USD")
			price := r.Ask + 0.3
			log.Println("---> Price:", price)
			amount := -1 * (SToF(data.ETH))
			log.Println("---> Amount:", amount)
			orderID, status = SubmitOrder(bfxPriv, price, amount)
		} else {
			log.Println("---> Buy ETH, Sell USD")
			price := r.Bid - 0.3
			log.Println("---> Price:", price)
			amount := SToF(data.USD) / price
			log.Println("---> Amount:", amount)
			orderID, status = SubmitOrder(bfxPriv, price, amount)
		}
	} else {
		log.Println("---> MOVING AVERAGE: No rebalance, HODL position")
	}
	/* 3. Update order status */
	if orderID != 0 {
		sheet.Update(data.Id, 12, Int64ToString(orderID))
		sheet.Update(data.Id, 13, status)
		sheet.Synchronize()
	} else {
		log.Println("---> No orderID received")
	}
	return 0
}

/*
==
Receives the price and the amount in order to submit an order in
Bitfinex and returns the orderID, which can be used to check
the status of the submitted order.
==
*/

func SubmitOrder(bfxPriv *rest.Client, price float64, amount float64) (int64, string) {
	response, err := bfxPriv.Orders.SubmitOrder(&order.NewRequest{
		Symbol:        "tETHUSD",
		CID:           time.Now().Unix() / 1000,
		Amount:        amount,
		Type:          "EXCHANGE LIMIT",
		Price:         price,
		AffiliateCode: "0xib78UF5",
	})
	if err != nil {
		panic(err)
	}
	orders := response.NotifyInfo.(*order.Snapshot)
	var orderID int64
	var status string
	for _, o := range orders.Snapshot {
		orderID = o.ID
		status = o.Status
	}
	log.Println("---> OrderID:", orderID)
	return orderID, status
}

/*
==
Checks the status of the order an updates the spreadsheet if an order
gets EXECUTED
==
*/

func MonitorOrderStatus(bfxPriv *rest.Client, sheet *spreadsheet.Sheet) {
	positions := QueryDB(sheet, "22:00:00")
	row := positions[len(positions)-1]

	/* Parameters */
	statusRow := row.Status
	ethUnits := SToF(row.ETH)
	id := int64(StringToInt(row.OrderID))

	order, err := bfxPriv.Orders.GetHistoryByOrderId(id)
	if err != nil {
		log.Println("--> orderID:", err)
	} else {
		if strings.Contains(statusRow, "EXECUTED") {
			/* Do nothing, the notification was already sent */
			log.Println("---> This order was EXECUTED and previously recorded")
		} else if statusRow == "ACTIVE" {
			if strings.Contains(order.Status, "EXECUTED") {
				log.Println("---> Order EXECUTED: updating spreadsheet")
				sheet.Update(row.Id, 13, order.Status)
				if ethUnits > 0 {
					/* Sell ETH, buy USD */
					sheet.Update(row.Id, 9, "0")
					sheet.Update(row.Id, 11, Float64ToString(-1*order.Price*order.AmountOrig*(1-0.001))) // Need testing
				} else {
					/* Buy ETH, sell USD */
					sheet.Update(row.Id, 9, Float64ToString(order.AmountOrig*(1-0.001)))
					sheet.Update(row.Id, 11, "0")
				}
				sheet.Synchronize()
			}
		}
		log.Println("---> Status from API:", order.Status)
	}
}

/*
==
Compute moving average. A "moving average" is the average of the last N values. A 20 day moving average is the average of the last 20 closing prices within a specific interval (i.e. hourly, daily)
==
*/
func ComputeMovingAverage(period int, positions []models.Position) {
	//1. Get all rows (in positions form) from spreadsheet
	//2. Identify those that are not filled
	//3. Compute MA
}
