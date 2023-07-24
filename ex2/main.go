package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var typ string
	if len(os.Args) <= 2 {
		fmt.Println("Format: <-type> <list of Int or String>")
		return
	}
	typ = os.Args[1]
	var argumentsInt []int
	var argumentsString []string

	for i := 2; i < len(os.Args); i++ {
		if typ == "-int" {
			val, err := strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Println("list contain an element which is not an Integer")
				return
			}
			argumentsInt = append(argumentsInt, val)
		} else if typ == "-string" {
			argumentsString = append(argumentsString, os.Args[i])
		} else {
			fmt.Println("Type undefinied")
		}

	}
	if len(argumentsInt) != 0 {
		fmt.Println("Before sort")
		fmt.Println(argumentsInt)
		sort.Ints(argumentsInt)
		fmt.Println("After sort")
		fmt.Println(argumentsInt)
	} else if len(argumentsString) != 0 {
		fmt.Println("Before sort")
		fmt.Println(argumentsString)
		sort.Strings(argumentsString)
		fmt.Println("After sort")
		fmt.Println(argumentsString)
	}

}
