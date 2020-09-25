/*
This program contains an algorithm manager for
a number of strategies to trade Bitcoin and Ether

V 1.0 20 DAY Moving Average, 09/21/2020
1. Choose where to fetch the prices (Source)
2. Make connection between Source and Google Spredsheets
3. Automate the store of data each hour (hourly price from Source)

4. Enter formulas for computing the Moving Average 20 day
5. Connect the exchange with the right API permissions to create/cancel orders
6. Set up BUY/SELL via API

*/
package main

import (
	"trading/client"
	"trading/spreadsheet"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func updateCandles() {
	_, bitfinex := client.ConnectionBitfinex()
	sh := client.ConnectionGoogle("1yLdidIUEIVJNVnSmMTKkALBj76cF8bI_HSGoR0QmFUg")
	sheet, _ := sh.SheetByTitle("ETH-20DMA")
	candles := spreadsheet.GetCandles(bitfinex, "tETHUSD", sheet)
	//trades := spreadsheet.ReadCsv()
	spreadsheet.WriteCandles(candles, sheet)
	//spreadsheet.MoveTrades(4, sheet, sheetSold)
}

func main() {
	updateCandles()
}
