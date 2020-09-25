package models

type Trade struct {
	Id       int
	Coin     string
	Base     string
	Exchange string
	Date     string
	Units    string
	BuyRate  string
	UUID     string
	BuyFee   string
	NetUnits string
	Invested string
	Rate     string
	Price    string
	SellFee  string
	NetSale  string
	Profit   string
	Status   string
}

/*
Error: if uncommented there seems to be an overlap between Trade
and Candle model.

Hint: maybe 'Id' can be used only once.

type ById []Trade

func (a ById) Len() int           { return len(a) }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
*/
