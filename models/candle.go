package models

type Candle struct {
	Id        int
	Asset     string
	Timestamp string
	Time      string
	Open      string
	Close     string
	High      string
	Low       string
	Volume    string
}

type ById []Candle

func (a ById) Len() int           { return len(a) }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
