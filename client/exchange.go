package client

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"os"
)

func ConnectionBitfinex() (client, bitfinex *rest.Client) {
	key := os.Getenv("BFX_KEY")
	secret := os.Getenv("BFX_SECRET")
	uri := "https://api.bitfinex.com/v2/"
	client = rest.NewClientWithURL(uri).Credentials(key, secret)
	bitfinex = rest.NewClient()
	return client, bitfinex
}
