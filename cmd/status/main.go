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

	timestamp := "22:00:00"

	sheetETH, _ := sh.SheetByTitle("ETH")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetETH, timestamp)

	sheetBTC, _ := sh.SheetByTitle("BTC")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetBTC, timestamp)

	sheetLTC, _ := sh.SheetByTitle("LTC")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetLTC, timestamp)

	sheetLINK, _ := sh.SheetByTitle("LINK")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetLINK, timestamp)

	sheetADA, _ := sh.SheetByTitle("ADA")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetADA, timestamp)

	sheetALGO, _ := sh.SheetByTitle("ALGO")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetALGO, timestamp)

	sheetATOM, _ := sh.SheetByTitle("ATOM")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetATOM, timestamp)

	sheetDOT, _ := sh.SheetByTitle("DOT")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetDOT, timestamp)
}
