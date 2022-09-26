package incident

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Using an empty struct{} here has the advantage that it doesn't require any additional space and Go's
// internal map type is optimized for that kind of values.
// Therefore, map[string] struct{} is a popular choice for sets in the Go world.
var validSortFields = map[string]struct{}{"discovered": {}, "status": {}}
var validSortDirections = map[string]struct{}{"ascending": {}, "descending": {}}
var validColumns = map[string]struct{}{"id": {}, "name": {}, "discovered": {}, "description": {}, "status": {}}

func ValidateFlags(incidentInput, sortDirection, sortField string) error {
	_, ok := validSortDirections[sortDirection]
	if !ok {
		return errors.New("Invalid sort direction")
	}
	_, ok = validSortFields[sortField]
	if !ok {
		return errors.New("Invalid sort field")
	}
	if incidentInput == "" {
		return errors.New("Input Json not provided")
	}

	return nil
}

func SortIncidents(incidents []Incident, sortDirection, sortField string) {
	direction := sortDirection == "ascending"

	switch sortField {
	case "status":
		sort.Slice(incidents, func(i, j int) bool {
			if direction {
				return incidentLessByDiscovered(incidents, i, j)
			}
			return !incidentLessByDiscovered(incidents, i, j)
		})

	case "discovered":
		sort.Slice(incidents, func(i, j int) bool {
			if direction {
				return incidentLessByDiscovered(incidents, i, j)
			}
			return !incidentLessByDiscovered(incidents, i, j)
		})
	}
}

func incidentLessByStatus(incident []Incident, i, j int) bool {
	return incident[i].Status < incident[j].Status
}

func incidentLessByDiscovered(incident []Incident, i, j int) bool {
	return incident[i].Discovered.String() < incident[j].Discovered.String()
}

func CreateCSVfromIncidents(path string, incidents []Incident, columns []string) {
	ouFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating output file, %v", err)
	}

	writer := csv.NewWriter(ouFile)

	defer ouFile.Close()
	defer writer.Flush()

	if err := writer.Write(columns); err != nil {
		log.Fatalf("Error while writing header in csv %v\n", err)
	}
	for _, row := range incidents {
		var csvRow []string
		for _, columnName := range columns {
			csvRow = append(csvRow, GetColumnValue(&row, columnName))
		}

		if err := writer.Write(csvRow); err != nil {
			log.Fatalf("Error writing row: %v\n", err)
		}
	}
}

func ValidateColumns(columns string) ([]string, error) {

	if columns == "" {
		// all valid comlumns
		return []string{"id", "name", "discovered", "description", "status"}, nil
	}

	sc := strings.Split(columns, ",")
	selectedColumns := []string{}

	for _, column := range sc {
		selectedColumns = append(selectedColumns, strings.Trim(column, " "))
	}

	for _, column := range selectedColumns {
		_, ok := validColumns[strings.Trim(column, " ")]
		if !ok {
			return nil, fmt.Errorf("Invalid column: %v, valida columns are: id, name, discovered, description, status", column)
		}
	}

	return selectedColumns, nil
}
