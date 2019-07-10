package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type record struct {
	Date     time.Time `json:"date"`
	Open     float64   `json:"open"`
	High     float64   `json:"high"`
	Low      float64   `json:"low"`
	Close    float64   `json:"close"`
	Volume   int64     `json:"volume"`
	AdjClose float64   `json:"adjClose"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	newFile, err := os.Create("index.html")

	if err != nil {
		log.Fatalln(err)
	}

	records := getRecords()
	err = tpl.Execute(newFile, records)

	if err != nil {
		log.Fatalln(err)
	}
}

func getRecords() []record {
	csvFile, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	i := 0

	var records []record

	for {
		line, err := reader.Read()

		if i == 0 {
			i++
			continue
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		date, _ := time.Parse("2006-01-02", line[0])
		open, _ := strconv.ParseFloat(line[1], 64)
		high, _ := strconv.ParseFloat(line[2], 64)
		low, _ := strconv.ParseFloat(line[3], 64)
		closeValue, _ := strconv.ParseFloat(line[4], 64)
		volume, _ := strconv.ParseInt(line[5], 10, 64)
		adjClose, _ := strconv.ParseFloat(line[6], 64)

		if err != nil {
			log.Fatal(err)
		}

		records = append(records, record{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    closeValue,
			Volume:   volume,
			AdjClose: adjClose,
		})
	}

	return records
}
