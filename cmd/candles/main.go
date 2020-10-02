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
	sh := client.ConnectionGoogle("1yLdidIUEIVJNVnSmMTKkALBj76cF8bI_HSGoR0QmFUg")
	sheet, _ := sh.SheetByTitle("ETH-20DMA")
	candles := spreadsheet.GetCandles(bfxPub, "tETHUSD", sheet)
	spreadsheet.WriteCandles(candles, sheet)
}
