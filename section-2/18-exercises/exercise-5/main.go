package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fmtDate": fmtDate,
}

type record struct {
	Date     time.Time `json:"date"`
	Open     float64   `json:"open"`
	High     float64   `json:"high"`
	Low      float64   `json:"low"`
	Close    float64   `json:"close"`
	Volume   int64     `json:"volume"`
	AdjClose float64   `json:"adjClose"`
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	records, error := readFile()
	if error != nil {
		log.Fatalln(error)
	}
	jsonRecords, error := json.Marshal(records)
	if error != nil {
		log.Fatalln(error)
	}
	jsonRecordsString := string(jsonRecords)
	http.HandleFunc("/data", func(response http.ResponseWriter, request *http.Request) {
		log.Println("Responding to request for /data")
		response.Header().Add("Content-Type", "application/json")
		fmt.Fprint(response, jsonRecordsString)
	})
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Println("Responding to request for /")
		response.Header().Add("Content-Type", "text/html")
		tpl.ExecuteTemplate(response, "index.gohtml", records)
	})
	log.Println("Starting Server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func fmtDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func readFile() ([]record, error) {
	file, error := os.Open("./table.csv")
	if error != nil {
		return []record{}, error
	}
	data, error := csv.NewReader(file).ReadAll()
	if error != nil {
		return []record{}, error
	}

	if len(data) <= 1 {
		return []record{}, nil
	}
	records := make([]record, len(data)-1)

	for index, row := range data[1:] {
		date, error := time.Parse("2006-01-02", row[0])
		if error != nil {
			return []record{}, error
		}
		open, error := strconv.ParseFloat(row[1], 64)
		if error != nil {
			return []record{}, error
		}
		high, error := strconv.ParseFloat(row[2], 64)
		if error != nil {
			return []record{}, error
		}
		low, error := strconv.ParseFloat(row[3], 64)
		if error != nil {
			return []record{}, error
		}
		close, error := strconv.ParseFloat(row[4], 64)
		if error != nil {
			return []record{}, error
		}
		volume, error := strconv.ParseInt(row[5], 10, 64)
		if error != nil {
			return []record{}, error
		}
		adjClose, error := strconv.ParseFloat(row[6], 64)
		if error != nil {
			return []record{}, error
		}

		records[index] = record{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			Volume:   volume,
			AdjClose: adjClose,
		}
	}

	// The first record is a header, so skip it
	return records, nil
}
