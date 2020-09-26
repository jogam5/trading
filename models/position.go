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
	ETH              string
	ETHValue         string
	USD              string
}
