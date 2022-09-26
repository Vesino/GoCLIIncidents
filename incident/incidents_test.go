package incident

import (
	"fmt"
	"testing"
	"time"
)

func TestValidFlags(t *testing.T) {
	inputJson, sortDirection, sortField := "", "invalid1", "invalid2"

	err := ValidateFlags(inputJson, sortDirection, sortField)
	if err == nil {
		t.Fatalf("ValidateFlags('', 'invalid1', 'invalid2') should return error")
	}
}

func TestValidateColumns(t *testing.T) {
	columns := "id, name, discovered"
	selectedColumns, _ := ValidateColumns(columns)

	exp := []string{"id", "name", "discovered"}
	if len(selectedColumns) != len(exp) {
		t.Fatalf("Len of selected columns: %v diff of %v", selectedColumns, exp)
	}

	invalidColumns := "id, name, invalid1"

	_, err := ValidateColumns(invalidColumns)
	expError := fmt.Errorf("Invalid column: %v, valida columns are: id, name, discovered, description, status", "invalid1")

	if expError.Error() != err.Error() {
		t.Fail()
	}
}

func TestSortIncidents(t *testing.T) {
	date1, _ := time.Parse("2006-01-02", "2018-04-02")
	date2, _ := time.Parse("2006-01-02", "2018-02-19")

	cases := []struct {
		desc                 string
		incidentA, IncidentB Incident
		fn                   func([]Incident, int, int) bool
		expected             bool
	}{
		{"name incidentA > incidentB",
			Incident{
				Id:          1,
				Name:        "Misdirected email",
				Discovered:  IncidentDiscovered(date1),
				Description: "Patient's medical records faxed to wrong number",
				Status:      "New"},
			Incident{
				Id:          2,
				Name:        "Misdirected fax",
				Discovered:  IncidentDiscovered(date2),
				Description: "Patient's medical records emailed to wrong email",
				Status:      "In Progress"},
			incidentLessByStatus, false},
	}
	for _, c := range cases {
		incidents := []Incident{c.incidentA, c.IncidentB}
		got := c.fn(incidents, 0, 1)
		if c.expected != got {
			t.Errorf("%s: expected %v, got %v", c.desc, c.expected, got)
		}
	}
}
