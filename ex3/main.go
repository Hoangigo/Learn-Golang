package main
import (
	"fmt"
	"github.com/Hoangigo/ex3/utils"
)
func main() {
	word := "a123bc34d8ef34"
	count := utils.NumDifferentIntegers(word)
	fmt.Println("Number of different Integers:", count)
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
	count2 := utils.CountRectangles(arr)
	fmt.Println("Number of Rectangles:", count2)
}