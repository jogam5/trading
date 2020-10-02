package client

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"os"
)

func ConnectionBitfinex() (bfxPriv, bfxPub *rest.Client) {
	key := os.Getenv("BFX_KEY")
	secret := os.Getenv("BFX_SECRET")
	uri := "https://api.bitfinex.com/v2/"
	bfxPriv = rest.NewClientWithURL(uri).Credentials(key, secret)
	bfxPub = rest.NewClient()
	return bfxPriv, bfxPub
}
