package pricedata

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadCSV reads CSV file for OHLCV stock price data and return it as a slice of 2-dimensional array
// CSV should be structured as follows:
//	- Date
//	- Open
//	- High
//	- Low
//	- Close
//	- Volume
//	- Adjusted Closing price
// "2018-01-29","4955","5080","4945","4965","742800","4965"
func ReadCSV(filepath string) ([][]int, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	fmt.Println(rows)
}
