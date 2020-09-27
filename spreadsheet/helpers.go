package spreadsheet

import (
	"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"strconv"
	"strings"
	"trading/models"
)

/*
==
Query the data base and return an slice of a struct Position. This
slice contains the specific information that will be used to feed
the algorithm.
==
*/

func QueryDB(sheet *spreadsheet.Sheet, timestamp string) []models.Position {
	/* Fetch candles */
	log.Println("Query data base")
	col := sheet.Columns[1]

	index := 1
	positions := []models.Position{}
	for _, v := range col {
		if strings.Contains(v.Value, timestamp) {
			//log.Println(sheet.Rows[v.Row][0].Value)
			p := models.Position{
				Id:               index,
				Timestamp:        sheet.Rows[v.Row][0].Value,
				Open:             sheet.Rows[v.Row][2].Value,
				MovingAverage:    sheet.Rows[v.Row][4].Value,
				PriceAboveMA:     SToBool(sheet.Rows[v.Row][5].Value),
				PriceCrossMA:     SToBool(sheet.Rows[v.Row][6].Value),
				Rebalance:        SToBool(sheet.Rows[v.Row][7].Value),
				PreviousPosition: sheet.Rows[v.Row][8].Value,
				ETH:              sheet.Rows[v.Row][9].Value,
				ETHValue:         sheet.Rows[v.Row][10].Value,
				USD:              sheet.Rows[v.Row][11].Value,
			}
			positions = append(positions, p)
			index = index + 1
		}
	}
	//log.Println(positions)
	return positions
}

func FetchPairs(sheet *spreadsheet.Sheet, exchange string) []models.Pair {
	/* Fetch pairs of coins */
	log.Println("Fetching data")
	col := sheet.Columns[0]

	/* Loop over pairs */
	pairs := []models.Pair{}
	for _, v := range col {
		if v.Value == "Coin" {
			continue
		} else if v.Value == "Closed" {
			break
		} else if len(v.Value) > 0 {
			base := sheet.Rows[v.Row][1]
			exc := sheet.Rows[v.Row][2]
			if exc.Value == exchange {
				p := models.Pair{
					Row:      int(v.Row),
					Coin:     v.Value,
					Base:     base.Value,
					Exchange: exc.Value,
				}
				pairs = append(pairs, p)
			}
		}
	}
	return pairs
}

func ReturnLastCell(colNumber uint, sheet *spreadsheet.Sheet) spreadsheet.Cell {
	/* Returns a last not null cell of a specific column */
	var last spreadsheet.Cell
	for _, cell := range sheet.Columns[colNumber] {
		if cell.Value == "" {
			break
		} else {
			last = cell
		}
	}
	return last
}

func ToI(s string) int {
	/* Converts string to int */
	i, _ := strconv.Atoi(s)
	return i
}

func ToF(s string) float64 {
	/* Converts string to float */
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

func FToS(num float64) string {
	/* Convert float64 to string */
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func Int64ToS(num int64) string {
	/* Convert int64 to string */
	return strconv.FormatInt(num, 10)
}

func SToBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}
