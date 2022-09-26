package incident

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
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

func GetColumnValue(i *Incident, column string) string {
	e := reflect.ValueOf(i).Elem()

	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Name == strings.Title(column) {
			return fmt.Sprint(e.Field(i).Interface())
		}
	}
	return ""
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
