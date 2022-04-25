package structag

import (
	"testing"
)

func TestDecode(t *testing.T) {
	src := map[string]string{
		"str":  "string data",
		"bool": "true",
	}
	var ms MapStruct
	Decode(&ms, src)

	if ms.Str != "string data" {
		t.Errorf("expect: %v, actually: %v", "string data", ms.Str)
	}
	if !ms.Bool {
		t.Errorf("expect: %v, actually: %v", true, ms.Bool)
	}
}
