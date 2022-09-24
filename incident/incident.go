package incident

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type Incident struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Discovered  IncidentDiscovered `json:"discovered"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
}

type IncidentDiscovered time.Time

func (i *IncidentDiscovered) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	*i = IncidentDiscovered(t)
	return nil
}

func (i IncidentDiscovered) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(i).Format("2006-01-02")), nil
}

func (i IncidentDiscovered) String() string {
	bteData, _ := i.MarshalJSON()
	return string(bteData)
}

func UnmarshallIncident(data []byte) []Incident {
	var incidents []Incident

	if err := json.Unmarshal(data, &incidents); err != nil {
		log.Fatalf("Error Unmarshall, the error was: %v", err)
	}
	return incidents
}
