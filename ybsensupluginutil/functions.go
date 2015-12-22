// Library for general purpose functions used by the Yieldbot Infrastructure
// teams in sensu.
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package ybsensupluginutil

import (
	"fmt"
	"os"
	"strings"
)

// EHndlr is for generic error handling in all Yieldbot monitoring packages.
func EHndlr(e error) {
	if e != nil {
		fmt.Printf("ERROR: %v", e)
		os.Exit(129)
	}
}

// Exit method for all sensu checks that will print the output and desired
// exit code
func Exit(args ...string) {
	// YELLOW need to make sure that condition exists
	var exitCode int
	output := ""

	if 1 > len(args) {
		panic("Not enough parameters.")
	}

	for i, p := range args {
		switch i {
		case 0: // name
			param, ok := p.(str)
			if !ok {
				panic("1st parameter not type string.")
			}

			for k := range MonitoringErrorCodes {
				if k == strings.ToUpper(param) {
					exitCode = MonitoringErrorCodes[k]
				}
			}

		case 1: // x
			param, ok := p.(string)
			if !ok {
				panic("2nd parameter not type string.")
			}
			output = param

		default:
			panic("Wrong parameter count.")
		}
	}

	fmt.Printf("%v", output)
	os.Exit(exitCode)
}

func CamelCaseAll(ss ...string) []string {

	camelCasedStrings := make([]string, len(ss))

	for i, s := range ss {
		camelCasedStrings[i] = stringcase.ToCamelCase(s)
	}

	return camelCasedStrings
}

func main() {
	ss1 := CamelCaseAll("hello world", "Apple Banana Cherry")
	// = []string{"helloWorld", "appleBananaCherry"}

	ss2 := CamelCaseAll("More than meets the eye")
	// = []string{"moreThanMeetsTheEye"}

	ss3 := CamelCaseAll()
	// = []string{}

	fmt.Printf("ss1 = %v\n", ss1)
	fmt.Printf("ss2 = %v\n", ss2)
	fmt.Printf("ss3 = %v\n", ss3)

	// Output:
	// ss1 = [helloWorld appleBananaCherry]
	// ss2 = [moreThanMeetsTheEye]
	// ss3 = []
}
