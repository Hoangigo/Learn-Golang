package tests

import (
	"testing"

	"github.com/Hoangigo/ex3/utils"
)

func TestNumberOfDifferentIntegers(t *testing.T) {
	word := "a123bc34d8ef34"
	expected := utils.NumDifferentIntegers(word)
	result := 3
	if result != expected {
		t.Errorf("TestNumberOfDifferentIntegers(word) returned %d, expected %d", result, expected)
	}
	word = "a123bc34d8ef345"
	expected = utils.NumDifferentIntegers(word)
	result = 4
	if result != expected {
		t.Errorf("TestNumberOfDifferentIntegers(word) returned %d, expected %d", result, expected)
	}
	word = "abcscsaccsc"
	expected = utils.NumDifferentIntegers(word)
	result = 0
	if result != expected {
		t.Errorf("TestNumberOfDifferentIntegers(word) returned %d, expected %d", result, expected)
	}
	word = "1111111111"
	expected = utils.NumDifferentIntegers(word)
	result = 1
	if result != expected {
		t.Errorf("TestNumberOfDifferentIntegers(word) returned %d, expected %d", result, expected)
	}
}
