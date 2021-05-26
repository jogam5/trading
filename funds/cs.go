/*
This file contains an algorithm manager for
the Computer Science Fund
*/

package funds

import (
	"log"
	"trading/client"
	"trading/models"
	"trading/spreadsheet"
)

/*
==
Manage the entrys and exits of the fund
==
*/

func RunCSFund() {
	bfxPriv, bfxPub := client.ConnectionBitfinex()
	log.Println(bfxPriv)
	sh := client.ConnectionGoogle("1Zp0jUHy5l2WdY6Je7_ltca7DVgPBE6VO8kDRtYxOcwo")

	timestamp := "19:00:00"
	assets := []models.Asset{
		models.Asset{Name: "ETH", QueryTag: "tETHUSD"},
		models.Asset{Name: "LTC", QueryTag: "tLTCUSD"},
		models.Asset{Name: "OMG", QueryTag: "tOMG:USD"},
		models.Asset{Name: "XTZ", QueryTag: "tXTZUSD"},
		models.Asset{Name: "BTC", QueryTag: "tBTCUSD"},
		models.Asset{Name: "LINK", QueryTag: "tLINK:USD"},
	}

	for _, v := range assets {
		log.Println(v.Name)
		sheet, _ := sh.SheetByTitle(v.Name)
		positions := spreadsheet.QueryDB(sheet, timestamp)
		spreadsheet.MovingAverage(sheet, bfxPriv, bfxPub, positions, v.QueryTag)
	}
}

//spreadsheet.SubmitOrder(bfxPriv, "tLINK:USD", 25.0, 0.2)
