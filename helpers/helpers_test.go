package helpers

import (
	"testing"
)

func TestInValidUrl(t *testing.T) {
	t.Log("Running : test url validity")

	testCasesInvalid := []string{
		"www.google.com",
		"google.com",
	}

	for _, each := range testCasesInvalid {
		if IsValidUrl(each) {
			t.Errorf("Failed: url validation: %s", each)
		}
	}
}

func TestValidUrl(t *testing.T) {
	t.Log("Running : test url validity")

	testCasesInvalid := []string{
		"http://www.google.com",
		"https://www.google.com",
	}

	for _, each := range testCasesInvalid {
		if !IsValidUrl(each) {
			t.Errorf("Failed: url validation: %s", each)
		}
	}

}
