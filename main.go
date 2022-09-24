package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var incidentInput = flag.String("json-input", "", "Jsong paylod which contains incident data")
var sortDirection = flag.String("sortdirection", "ascending", "Sort columns in the specified direction, optional values: ascending or descending")
var sortField = flag.String("sortfield", "discovered", "Sort columns by field, could, optional values: discovered or status")

type Incident struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Discovered  IncidentDiscovered `json:"discovered"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
}

type IncidentDiscovered time.Time

func (i *IncidentDiscovered) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*i = IncidentDiscovered(t) //set result using the pointer
	return nil
}

func (i IncidentDiscovered) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(i).Format("2006-01-02")), nil
}

func (i IncidentDiscovered) String() string {
	bteData, _ := i.MarshalJSON()
	return string(bteData)
}

func main() {
	flag.Parse()

	if *incidentInput == "" {
		log.Fatal("Incident Input not provided")
	}

	data := []byte(*incidentInput)
	var incidents []Incident
	if err := json.Unmarshal(data, &incidents); err != nil {
		log.Fatalf("Error Unmarshall, the error was: %v", err)
	}

	sortIncidents(incidents, *sortDirection, *sortField)

	ouFile, err := os.Create("test.csv")
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

func sortIncidents(incidents []Incident, sortDirection, sortField string) {

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
