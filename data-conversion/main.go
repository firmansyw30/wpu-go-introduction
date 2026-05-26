package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Integer to String
	var num int = 42
	var str string = strconv.Itoa(num)
	fmt.Println("Integer to String:", str)

	// String to Integer (Required error handling)
	var strNum string = "123"
	var intNum, err = strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String to Integer:", intNum)
	}

	// Another String to Integer
	var anotherStrNum string = "456"
	var anotherIntNum, _ = strconv.Atoi(anotherStrNum)
	fmt.Println("Another String to Integer:", anotherIntNum)

	// Boolean to string
	truth := true
	boolToStr := strconv.FormatBool(truth)
	fmt.Println("Boolean to String:", boolToStr)

	// String to boolean
	opposite := "false" // only true or false value
	strToBool, err := strconv.ParseBool(opposite)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String to Boolean:", strToBool)
	}

	// Another string to boolean
	anotherStrToBool := "true"
	strToBool2, _ := strconv.ParseBool(anotherStrToBool) //strconv.ParseBool required at least 2 arguments for parsing
	fmt.Println("Another String to Boolean:", strToBool2)

}
