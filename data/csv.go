package data

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ODD = "odd"
const EVEN = "even"

type CSVData struct {
	filePath string
}

func NewCSVData(filePath string) CSVData {
	return CSVData{filePath: filePath}
}

func (cs CSVData) GetRecordById(id int, rowType string) ([]string, error) {

	file, err := cs.getFile()
	if err != nil {
		return nil, err
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

func (cs CSVData) WorkerReader(csvLine chan<- []string, items int, typeNumber string) {
	file, err := cs.getFile()
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineData := strings.Split(line, ",")
		id, _ := strconv.Atoi(lineData[0])

		if typeNumber == EVEN && id%2 == 0 {
			counter++
			if counter <= items {
				csvLine <- lineData
			} else {
				fmt.Println("*** items limit ***")
				break
			}
		} else if typeNumber == ODD && id%2 == 1 {
			counter++
			if counter <= items {
				csvLine <- lineData
			} else {
				fmt.Println("*** items limit ***")
				break
			}
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	close(csvLine)
}

func (cs CSVData) getFile() (*os.File, error) {
	file, err := os.Open(cs.filePath)
	if err != nil {
		return nil, errors.New("csv not found")
	}
	return file, nil
}
