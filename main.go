package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type BinData struct {
	CompanyName      string
	ICA              string
	AccountRangeFrom int
	AccountRangeTo   int
	BrandProductCode string
	BrandProductName string
	AcceptanceBrand  string
	Country          string
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Printf("open file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("read file: %v", err)
		return
	}

	if len(rows) < 2 {
		fmt.Println("rows < 2")
		return
	}

	var records []BinData
	for _, row := range rows[1:] {
		if len(row) < 8 {
			fmt.Println("incomplete line:", row)
			continue
		}
		low, err1 := strconv.Atoi(row[2])
		high, err2 := strconv.Atoi(row[3])
		if err1 != nil || err2 != nil {
			fmt.Printf("convert bin range failed: %s: %v", row, err)
			continue
		}
		records = append(records, BinData{
			CompanyName:      row[0],
			ICA:              row[1],
			AccountRangeFrom: low,
			AccountRangeTo:   high,
			BrandProductCode: row[4],
			BrandProductName: row[5],
			AcceptanceBrand:  row[6],
			Country:          row[7],
		})
	}

	var searchBIN int
	fmt.Print("BIN(10): ")
	fmt.Scan(&searchBIN)

	for _, record := range records {
		if searchBIN >= record.AccountRangeFrom && searchBIN <= record.AccountRangeTo {
			b, _ := json.MarshalIndent(record, "", " ")
			fmt.Printf("%+v\n", string(b))
		}
	}
}
