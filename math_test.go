package main 

import "testing"

func SumTest(t *testing.T) {
	result := Sum(3, 4)
	expected := 7
	if result != expected {
		t.Errorf("Sum(3, 4) = %d; want %d", result, expected)
	}
}