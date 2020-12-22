package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testInputs := []struct {
		name      string
		input     []int
		expectecd []int
	}{
		{
			"No Duplicates",
			[]int{10, 9, -8, 2, 4, 1, 12, 56, 0, 5},
			[]int{-8, 0, 1, 2, 4, 5, 9, 10, 12, 56},
		},
		{
			"With Duplicates",
			[]int{1, 3, 0, 0, -1, 4},
			[]int{-1, 0, 0, 1, 3, 4},
		},
	}

	for _, test := range testInputs {
		t.Run(test.name, func(t *testing.T) {
			SelectionSort(test.input)
			if !reflect.DeepEqual(test.input, test.expectecd) {
				t.Errorf("Error while sorting, expecting: %v, found: %v", test.input, test.expectecd)
			}
		})
	}
}
