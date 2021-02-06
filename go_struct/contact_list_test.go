package main

import (
	"testing"
)

func TestContact(t *testing.T) {
	contact1 := contact{
		firstName: "Qobiljonov",
		lastName:  "Davrbek",
		number:    938554784,
	}

	if contact1.lastName != "Davrbek" {
		t.Errorf("Error lastName")
	}
}
