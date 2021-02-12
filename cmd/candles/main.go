/*
This program contains a function that is executed via a cron job.
Every specific time the program retrieves candle data from Bitfinex.
*/

package main

import (
	"trading/client"
	"trading/spreadsheet"
)

func main() {
	_, bfxPub := client.ConnectionBitfinex()
	sh := client.ConnectionGoogle("1MK6SUfDrVHQXWL7pUZzS3yxkWuIDecAvHqxXpSKHWL8")
	sheetETH, _ := sh.SheetByTitle("ETH")
	candlesETH := spreadsheet.GetCandles(bfxPub, "tETHUSD", sheetETH)
	spreadsheet.WriteCandles(candlesETH, sheetETH)

	sheetBTC, _ := sh.SheetByTitle("BTC")
	candlesBTC := spreadsheet.GetCandles(bfxPub, "tBTCUSD", sheetBTC)
	spreadsheet.WriteCandles(candlesBTC, sheetBTC)

	sheetLTC, _ := sh.SheetByTitle("LTC")
	candlesLTC := spreadsheet.GetCandles(bfxPub, "tLTCUSD", sheetLTC)
	spreadsheet.WriteCandles(candlesLTC, sheetLTC)

	sheetLINK, _ := sh.SheetByTitle("LINK")
	candlesLINK := spreadsheet.GetCandles(bfxPub, "tLINK:USD", sheetLINK)
	spreadsheet.WriteCandles(candlesLINK, sheetLINK)

}
