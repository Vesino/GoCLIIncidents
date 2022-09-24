package incident

import (
	"testing"
)

func TestValidFlags(t *testing.T) {
	inputJson, sortDirection, sortField := "", "invalid1", "invalid2"

	err := ValidateFlags(inputJson, sortDirection, sortField)
	if err == nil {
		t.Fatalf("ValidateFlags('', 'invalid1', 'invalid2') should return error")
	}
}
