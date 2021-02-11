package main

import "testing"

func TestTast(t *testing.T) {
	task1 := task{
		title: "Fibonacci",
		body:  "You must find n th fibonacci number",
	}

	if task1.title != "Fibonacci" {
		t.Errorf("test.go is not work! wrong title")
	}
}
