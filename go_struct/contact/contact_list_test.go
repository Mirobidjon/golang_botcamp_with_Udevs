package main

import "testing"

func TestContact(t *testing.T) {
	contact1 := contact{
		firstName: "Kamolov",
		lastName:  "Xusniddin",
		number:    978851887,
	}

	if contact1.fistName != "Kamolov" {
		t.Errorf("test.go is not work! wrong title")
	}
}
