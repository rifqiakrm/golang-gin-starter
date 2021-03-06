package utils

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

type M [][]string

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("can't close file while reading csv:", err)
		}
	}()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func ReadExcelFile(filePath string) [][]string {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Println("can't close file while reading excel:", err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal("Unable to parse file as XLSX for "+filePath, err)
	}

	return rows
}
