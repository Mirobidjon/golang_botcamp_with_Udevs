package main

import "testing"

func TestContact(t *testing.T) {
	contact1 := contact{
		firstName: "Kamolov",
		lastName:  "Xusniddin",
		number:    978851887,
	}

	if contact1.firstName != "Kamolov" {
		t.Errorf("Contact struct don't work! wrong Name")
	}
}
