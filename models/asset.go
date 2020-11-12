package models

type Asset struct {
	Name      string
	Pair      string
	Candles   []Candle
	Positions []Position
	DateTime  string
}
