package utils

import "testing"

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	if uuid.String() == "" {
		t.Error("UUID could not be generated")
	}
}
