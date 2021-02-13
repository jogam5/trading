package models

type Position struct {
	Id               int
	Timestamp        string
	Open             string
	MovingAverage    string
	PriceAboveMA     bool
	PriceCrossMA     bool
	Rebalance        bool
	PreviousPosition string
	CoinUnits        string
	USDValue         string
	USDUnits         string
	OrderID          string
	Status           string
}

type PositionById []Position

func (a PositionById) Len() int           { return len(a) }
func (a PositionById) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a PositionById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
