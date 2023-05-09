package main

import "testing"

func TestCombinedNumber01(t *testing.T) {
	x := []int{1}
	actual := combinedNumber(x)
	expected := "1"
	if actual != expected {
		t.Error("error")
	}
}
