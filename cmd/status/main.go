/*
This program contains a function that is executed via a cron job. Every specific time the program monitors if an order has been executed and updates the spreadsheet accordingly.
*/

package main

import (
	"log"
	"trading/client"
	"trading/models"
	"trading/spreadsheet"
)

func main() {
	sh := client.ConnectionGoogle("1MK6SUfDrVHQXWL7pUZzS3yxkWuIDecAvHqxXpSKHWL8")
	bfxPriv, _ := client.ConnectionBitfinex()

	timestamp := "19:00:00"
	assets := []models.Asset{
		models.Asset{Name: "ETH"},
		models.Asset{Name: "BTC"},
		models.Asset{Name: "LTC"},
		models.Asset{Name: "LINK"},
		models.Asset{Name: "ADA"},
		models.Asset{Name: "ALGO"},
		models.Asset{Name: "ATOM"},
		models.Asset{Name: "DOT"},
		models.Asset{Name: "XRP"},
	}

	for _, v := range assets {
		log.Println(v.Name)
		sheet, _ := sh.SheetByTitle(v.Name)
		spreadsheet.MonitorOrderStatus(bfxPriv, sheet, timestamp)
	}

	sheetETHFund, _ := sh.SheetByTitle("ETH-FUND")
	spreadsheet.MonitorOrderStatus(bfxPriv, sheetETHFund, timestamp)

}
