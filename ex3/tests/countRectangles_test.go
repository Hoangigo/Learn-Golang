package tests

import (
	"testing"

	"github.com/Hoangigo/ex3/utils"
)

func TestCount(t *testing.T) {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}
	result := utils.CountRectangles(arr)
	expected := 6

	if result != expected {
		t.Errorf("CountRectangles(arr) returned %d, expected %d", result, expected)
	}
	arr = [][]int{
		{1, 0, 0, 1, 1, 0, 1},
		{0, 0, 0, 1, 1, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 0, 0},
		{1, 0, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 1},
	}
	result = utils.CountRectangles(arr)
	expected = 8

	if result != expected {
		t.Errorf("CountRectangles(arr) returned %d, expected %d", result, expected)
	}
	arr = [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	result = utils.CountRectangles(arr)
	expected = 1

	if result != expected {
		t.Errorf("CountRectangles(arr) returned %d, expected %d", result, expected)
	}
}
