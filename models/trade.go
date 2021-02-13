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
