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
	"trading/models"
	"trading/spreadsheet"
)

func main() {
	bfxPriv, bfxPub := client.ConnectionBitfinex()
	log.Println(bfxPriv)
	sh := client.ConnectionGoogle("1MK6SUfDrVHQXWL7pUZzS3yxkWuIDecAvHqxXpSKHWL8")

	timestamp := "19:00:00"
	assets := []models.Asset{
		models.Asset{Name: "ETH", QueryTag: "tETHUSD"},
		models.Asset{Name: "LTC", QueryTag: "tLTCUSD"},
		models.Asset{Name: "LINK", QueryTag: "tLINK:USD"},
		models.Asset{Name: "ALGO", QueryTag: "tALGUSD"},
		models.Asset{Name: "ATOM", QueryTag: "tATOUSD"},
		models.Asset{Name: "DOT", QueryTag: "tDOTUSD"},
		models.Asset{Name: "XRP", QueryTag: "tXRPUSD"},
	}

	for _, v := range assets {
		log.Println(v.Name)
		sheet, _ := sh.SheetByTitle(v.Name)
		positions := spreadsheet.QueryDB(sheet, timestamp)
		spreadsheet.MovingAverage(sheet, bfxPriv, bfxPub, positions, v.QueryTag)
	}

	//spreadsheet.SubmitOrder(bfxPriv, "tLINK:USD", 25.0, 0.2)
}
