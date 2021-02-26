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

	sheetADA, _ := sh.SheetByTitle("ADA")
	candlesADA := spreadsheet.GetCandles(bfxPub, "tADAUSD", sheetADA)
	spreadsheet.WriteCandles(candlesADA, sheetADA)

	sheetALGO, _ := sh.SheetByTitle("ALGO")
	candlesALGO := spreadsheet.GetCandles(bfxPub, "tALGUSD", sheetALGO)
	spreadsheet.WriteCandles(candlesALGO, sheetALGO)

	sheetATOM, _ := sh.SheetByTitle("ATOM")
	candlesATOM := spreadsheet.GetCandles(bfxPub, "tATOUSD", sheetATOM)
	spreadsheet.WriteCandles(candlesATOM, sheetATOM)

	sheetDOT, _ := sh.SheetByTitle("DOT")
	candlesDOT := spreadsheet.GetCandles(bfxPub, "tDOTUSD", sheetDOT)
	spreadsheet.WriteCandles(candlesDOT, sheetDOT)

	sheetEOS, _ := sh.SheetByTitle("EOS")
	candlesEOS := spreadsheet.GetCandles(bfxPub, "tEOSUSD", sheetEOS)
	spreadsheet.WriteCandles(candlesEOS, sheetEOS)

	sheetETC, _ := sh.SheetByTitle("ETC")
	candlesETC := spreadsheet.GetCandles(bfxPub, "tETCUSD", sheetETC)
	spreadsheet.WriteCandles(candlesETC, sheetETC)
}
