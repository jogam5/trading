/*
This program contains an algorithm manager for
a number of strategies to trade Bitcoin and Ether

**********************
20 DAY Moving Average
*********************


README:
https://www.makeareadme.com
*/

package main

import (
	"log"
	"trading/client"
	"trading/spreadsheet"
)

func main() {
	bfxPriv, bfxPub := client.ConnectionBitfinex()
	log.Println(bfxPriv)
	sh := client.ConnectionGoogle("1MK6SUfDrVHQXWL7pUZzS3yxkWuIDecAvHqxXpSKHWL8")

	/* ETH */
	sheetETH, _ := sh.SheetByTitle("ETH")
	positionsETH := spreadsheet.QueryDB(sheetETH, "18:00:00")
	spreadsheet.MovingAverage(sheetETH, bfxPriv, bfxPub, positionsETH, "tETHUSD")

	/* BTC */
	/*sheetBTC, _ := sh.SheetByTitle("BTC")
	positionsBTC := spreadsheet.QueryDB(sheetBTC, "18:00:00")
	spreadsheet.MovingAverage(sheetBTC, bfxPriv, bfxPub, positionsBTC, "tBTCUSD")
	*/

	/*sheetLTC, _ := sh.SheetByTitle("LTC")
	positionsLTC := spreadsheet.QueryDB(sheetLTC, "18:00:00")
	spreadsheet.MovingAverage(sheetLTC, bfxPriv, bfxPub, positionsLTC, "tLTCUSD")
	*/

	/*sheetLINK, _ := sh.SheetByTitle("LINK")
	positionsLINK := spreadsheet.QueryDB(sheetLINK, "18:00:00")
	spreadsheet.MovingAverage(sheetLINK, bfxPriv, bfxPub, positionsLINK, "tLINKUSD")
	*/

	//spreadsheet.SubmitOrder(bfxPriv, "tLINK:USD", 25.0, 0.2)
}
