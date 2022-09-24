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

	sort.Slice(incidents, func(i, j int) bool {
		comparition := strings.Compare(incidents[i].Discovered.String(), incidents[j].Discovered.String())
		if sortField == "status" {
			comparition = strings.Compare(incidents[i].Status, incidents[j].Status)
		}
		switch comparition {
		case -1:
			return !(sortDirection == "ascending")
		case 1:
			return sortDirection == "ascending"
		}
		return incidents[i].Status > incidents[j].Status
	})
}

func CreateCSVfromIncidents(path string, incidents []Incident) {
	ouFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating output file, %v", err)
	}

	writer := csv.NewWriter(ouFile)

	defer ouFile.Close()
	defer writer.Flush()

	if err := writer.Write([]string{"id", "name", "discovered", "description", "status"}); err != nil {
		log.Fatalf("Error while writing header in csv %v\n", err)
	}
	for _, row := range incidents {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(row.Id), row.Name, row.Discovered.String(), row.Description, row.Status)
		if err := writer.Write(csvRow); err != nil {
			log.Fatalf("Error writing row: %v\n", err)
		}
	}
}
