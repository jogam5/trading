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

	positions := []models.Position{}
	for _, v := range col {
		if strings.Contains(v.Value, timestamp) {
			//log.Println(sheet.Rows[v.Row][0].Value)
			p := models.Position{
				Id:               int(v.Row), /*casting uint into int*/
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
				OrderID:          sheet.Rows[v.Row][12].Value,
				Status:           sheet.Rows[v.Row][13].Value,
			}
			positions = append(positions, p)
		}
	}
	//log.Println(positions)
	return positions
}

/*
==
Returns a last not null cell of a specific column
==
*/

func ReturnLastCell(colNumber uint, sheet *spreadsheet.Sheet) spreadsheet.Cell {
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

/*
==
Find a specific value in the spreadsheet and return a cell
==
*/

func FindValue(sheet *spreadsheet.Sheet, value string) spreadsheet.Cell {
	var cellFound spreadsheet.Cell
	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Value == value {
				cellFound = cell
			}
		}
	}
	return cellFound
}

func SToI(s string) int {
	/* Converts string to int */
	i, _ := strconv.Atoi(s)
	return i
}

func SToF(s string) float64 {
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
