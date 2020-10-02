/*
This program contains a function that is executed via a cron job. Every specific time the program monitors if an order hasbeen executed and updates the spreadsheet accordingly.
*/

package main

import (
	"trading/client"
	"trading/spreadsheet"
)

func main() {
	sh := client.ConnectionGoogle("1yLdidIUEIVJNVnSmMTKkALBj76cF8bI_HSGoR0QmFUg")
	bfxPriv, _ := client.ConnectionBitfinex()
	sheet, _ := sh.SheetByTitle("ETH-20DMA")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheet)
}
