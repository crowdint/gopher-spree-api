package domain

import (
	"encoding/json"
	"testing"
)

func AssertEqualJson(t *testing.T, got interface{}, expected string) {
	jsonData, err := json.Marshal(got)
	if err != nil {
		t.Error("Marshaling error:", err.Error())
	}

	jsonString := string(jsonData)

	if jsonString != expected {
		t.Errorf("Mismacth: %s : %s", jsonString, expected)
	}
}
