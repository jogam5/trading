/*
This program contains a function that is executed via a cron job. Every specific time the program monitors if an order has been executed and updates the spreadsheet accordingly.
*/

package main

import (
	"trading/client"
	"trading/spreadsheet"
)

func main() {
	sh := client.ConnectionGoogle("1MK6SUfDrVHQXWL7pUZzS3yxkWuIDecAvHqxXpSKHWL8")
	bfxPriv, _ := client.ConnectionBitfinex()

	sheetETH, _ := sh.SheetByTitle("ETH")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetETH)

	sheetBTC, _ := sh.SheetByTitle("BTC")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetBTC)

	sheetLTC, _ := sh.SheetByTitle("LTC")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetLTC)

	sheetLINK, _ := sh.SheetByTitle("LINK")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetLINK)

}
