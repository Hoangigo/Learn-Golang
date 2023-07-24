package main

import (
	"fmt"
	"os"
)

func main() {
	var fName, lName, code string
	if len(os.Args) <= 2 {
		fmt.Println("Format: <first_name> <last_name> <code>")
		return
	}
	fName = os.Args[1]
	lName = os.Args[2]
	code = os.Args[3]
	tempCode := formatCode(code)
	fullName := formatName(fName, lName, tempCode)
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
func formatName(fName, lName, nameFormat string) string {
	if nameFormat == "first_last" {
		return fName + " " + lName
	} else if nameFormat == "last_first" {
		return lName + " " + fName
	} else {
		return "Undefined format"
	}
}
