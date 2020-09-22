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

type ById []Trade

func (a ById) Len() int           { return len(a) }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
