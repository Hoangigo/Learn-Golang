package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
				fmt.Println("Invalid input")
				return
			}
			argumentsInt = append(argumentsInt, val)
		} else if typ == "-string" {
			argumentsString = append(argumentsString, os.Args[i])
		} else {
			fmt.Println("Type undefinied")
			return
		}

	}
	if len(argumentsInt) != 0 {
		sort.Ints(argumentsInt)
		var strInts []string
		for _, val := range argumentsInt {
			strInts = append(strInts, strconv.Itoa(val))
		}
		fmt.Println("Output:", strings.Join(strInts, " "))	
		} else if len(argumentsString) != 0 {		
		sort.Strings(argumentsString)
		fmt.Println("Output:",strings.Join(argumentsString, " "))
	}

}
