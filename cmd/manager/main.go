/*
This program contains an algorithm manager for
a number of strategies to trade Bitcoin and Ether

**********************
20 DAY Moving Average
*********************

### V 1.0 - 09/21/2020
1. Choose where to fetch the prices (Source)
2. Make connection between Source and Google Spredsheets
3. Automate the store of data each hour (hourly price from Source)
4. Enter formulas for computing the Moving Average 20 day

### V 1.1 - 09/25/2020
5. Read actionable data from Spreadsheet
6. Connect the exchange with the API permissions to create orders
7. Set up BUY/SELL via API
8. Improve README

### V 1.2 - 09/29/2020
9. Write orderID in the Spreadsheet and monitors its status on Bitfinex
10. Update holdings
11. Reorganize code in CMD
12. Monitor via cronjob in an hourly basis if an order has been fulfilled
HERE:
13. Create blank spreadsheet for the project
14. Create a new Btfinex account for the project
17. Reorganize spreadsheet to reflect current value of holdings
16. Test automation of MA algorithm
15. Send notification when the Rebalance has taken place (SMS, Email, etc)

### V1.x
1x. Build a simple website to show how the strategy performs against Buy and Hold.
1x. Instead of rebalancing to USD, use CUSD (the Compound's version of USD) so that an interest can be accrued on the amount of USD.
1x. Instead of rebalancing to ETH, find a way to accrue some interest in Compound or AAVE.
1x. Testing the algorithm on other altcoins.
1x. Generalize the algorithm so that it can be used for other periods of the Moving Average (i.e. 1 Week Moving Average for long term strategies in Bitcoin)
1x. Use a Database instead of relying on Google Spreadsheet to store the data.

README:
https://www.makeareadme.com
*/

package main

import (
	"trading/client"
	"trading/spreadsheet"
)

func main() {
	bfxPriv, bfxPub := client.ConnectionBitfinex()
	sh := client.ConnectionGoogle("1yLdidIUEIVJNVnSmMTKkALBj76cF8bI_HSGoR0QmFUg")
	sheet, _ := sh.SheetByTitle("ETH-20DMA")
	positions := spreadsheet.QueryDB(sheet, "22:00:00")
	spreadsheet.MovingAverage(sheet, bfxPriv, bfxPub, positions)
}