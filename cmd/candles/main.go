/*
This program contains a function that is executed via a cron job.
Every specific time the program retrieves candle data from Bitfinex.

Improvements:
1. Update only at market open
*/

package main

import (
	"context"
	"fmt"
	"log"
	"trading/client"
	"trading/models"
	"trading/spreadsheet"

	//"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	conn, _ := pgxpool.Connect(context.Background(), "postgres://postgres:123@localhost:5432/trading_engine")
	fmt.Println(conn)
	defer conn.Close()

	_, bfxPub := client.ConnectionBitfinex()

	log.Println("### -> Regular Fund")
	sh := client.ConnectionGoogle("1MK6SUfDrVHQXWL7pUZzS3yxkWuIDecAvHqxXpSKHWL8")

	assets := []models.Asset{
		models.Asset{Name: "ETH", QueryTag: "tETHUSD"},
		/*models.Asset{Name: "LTC", QueryTag: "tLTCUSD"},
		models.Asset{Name: "LINK", QueryTag: "tLINK:USD"},
		models.Asset{Name: "ALGO", QueryTag: "tALGUSD"},
		models.Asset{Name: "ATOM", QueryTag: "tATOUSD"},
		models.Asset{Name: "DOT", QueryTag: "tDOTUSD"},
		models.Asset{Name: "XRP", QueryTag: "tXRPUSD"},
		models.Asset{Name: "BTC", QueryTag: "tBTCUSD"},
		models.Asset{Name: "ADA", QueryTag: "tADAUSD"},
		models.Asset{Name: "UNI", QueryTag: "tUNIUSD"},
		*/
	}

	/* Information is stored here */
	for _, v := range assets {
		log.Println(v.Name)
		sheet, _ := sh.SheetByTitle(v.Name)
		candles := spreadsheet.GetCandles(bfxPub, v.QueryTag, sheet)

		spreadsheet.WriteCandles(candles, sheet)
	}

	/* CS Fund */
	/* This fund is not working right now. The assets can be changed though.*/
	log.Println("### -> CS Fund")
	sh = client.ConnectionGoogle("1Zp0jUHy5l2WdY6Je7_ltca7DVgPBE6VO8kDRtYxOcwo")

	assets = []models.Asset{
		models.Asset{Name: "ETH", QueryTag: "tETHUSD"},
		//models.Asset{Name: "LTC", QueryTag: "tLTCUSD"},
		//models.Asset{Name: "OMG", QueryTag: "tOMGUSD"},
		//models.Asset{Name: "XTZ", QueryTag: "tXTZUSD"},
		//models.Asset{Name: "BTC", QueryTag: "tBTCUSD"},
		//models.Asset{Name: "LINK", QueryTag: "tLINK:USD"},
	}

	for _, v := range assets {
		log.Println(v.Name)
		sheet, _ := sh.SheetByTitle(v.Name)
		candles := spreadsheet.GetCandles(bfxPub, v.QueryTag, sheet)
		spreadsheet.WriteCandles(candles, sheet)
		spreadsheet.InsertCandles(v.Name, candles, conn)
	}
}
