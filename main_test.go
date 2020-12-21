package main

import (
	"io/ioutil"
	"testing"
)

func TestAddTwo(t *testing.T) {
	if addTwo(2) != 4 {
		t.Error("Expected 2 + 2 = 4")
	}
}

func TestTaleAddTwo(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expected int
	}{
		{"premier", 1, 3},
		{"second", -5, -3},
		{"troisieme", 12, 14},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := addTwo(test.input); output != test.expected {
				t.Fatalf("Test failed: %d inputed, %d outputed, expected: %d", test.input, output, test.expected)
			}
		})
	}
}

func TestData(t *testing.T) {
	filename := "testData/test.data"
	expectedString := "test data"

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Errorf("Could not red file: %v", filename)
	}

	readString := string(data)
	if readString != expectedString {
		t.Errorf("wanted: %v, read: %v", expectedString, readString)
	}
}

func TestScanUserInput(t *testing.T) {
	scanUserInput()
}
