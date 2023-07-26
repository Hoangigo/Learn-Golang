package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var fName, lName, code string
	if len(os.Args) <= 3 {
		fmt.Println("Format: <first_name> <middle_name...> <last_name> <code>")
		return
	}
	fName = os.Args[1]
	lName = os.Args[len(os.Args)-2]
	code = os.Args[len(os.Args)-1]
	code = strings.ToLower(code)
	middleNames := make([]string, 0)
	for i := 2; i < len(os.Args)-2; i++ {
		middleNames = append(middleNames, strings.ToLower(os.Args[i]))
	}
	tempCode := formatCode(code)
	fullName := formatName(fName, lName, tempCode, middleNames)
	fmt.Println(fullName)

}
func formatCode(code string) string {
	switch code {
	case "us":
		return "first_last"
	case "vn":
		return "last_first"
	default:
		return "undefined"
	}
}
func formatName(fName, lName, nameFormat string, middle_name[] string) string {
	if nameFormat == "first_last" {
		return fName + " "+ strings.Join(middle_name, " ")+ " " + lName
	} else if nameFormat == "last_first" {
		return lName + " "+ strings.Join(middle_name, " ")+ " " + fName
	} else {
		return "Undefined format"
	}
}
