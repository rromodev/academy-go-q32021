package data

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type CSVData struct {
	filePath string
}

func NewCSVData(filePath string) CSVData {
	return CSVData{filePath: filePath}
}

func (cs CSVData) GetRecordById(id int, rowType string) ([]string, error) {

	file, err := os.Open(cs.filePath)
	if err != nil {
		return nil, errors.New("csv not found")
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var csvRecord []string

	for _, record := range records {
		idRecord, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, errors.New("csv id not found")
		}
		if id != 0 && id == idRecord && rowType == record[1] {
			csvRecord = record
		}
	}
	if len(csvRecord) == 0 {
		return nil, errors.New("record not found")
	}
	return csvRecord, nil
}

func (cs CSVData) WriteRecord(record []string) (string, error) {
	file, err := os.OpenFile(cs.filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return "", errors.New("csv not found")
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write(record); err != nil {
		return "", errors.New("error writing record to file")
	}
	return "ok", nil
}
