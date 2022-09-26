package main

import (
	"flag"
	"log"

	"github.com/Vesino/GoCLIIncidents/incident"
)

var columns string

var incidentInput = flag.String("json-input", "", "JSON payload which contains incident data")
var sortDirection = flag.String("sortdirection", "ascending", "Sort columns in the specified direction, optional values: ascending or descending")
var sortField = flag.String("sortfield", "discovered", "Sort columns by field, could, optional values: discovered or status")
var csvPath = flag.String("path", "test.csv", "path to store the .csv file generated")

func main() {
	flag.StringVar(&columns, "columns", "", "Columns to output in CSV")
	flag.Parse()

	// validate flags options
	if err := incident.ValidateFlags(*incidentInput, *sortDirection, *sortField); err != nil {
		log.Fatalf("Invalid flag: %v", err)
	}

	selectedColumns, err := incident.ValidateColumns(columns)
	if err != nil {
		log.Fatalf("Column selected error: %v", err)
	}

	data := []byte(*incidentInput)

	incidents := incident.UnmarshallIncident(data)
	incident.SortIncidents(incidents, *sortDirection, *sortField)
	incident.CreateCSVfromIncidents(*csvPath, incidents, selectedColumns)
}
