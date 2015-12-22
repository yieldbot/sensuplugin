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
func Exit(s string, o string) {
	// YELLOW need to make sure that condition exists
	var exitCode int

	for k := range ybsensupluginutil.MonitoringErrorCodes {
		if k == strings.ToUpper(s) {
			exitCode = ybsensupluginutil.MonitoringErrorCodes[k]
		}
		fmt.Printf("%v", o)
		os.Exit(exitCode)
	}
}
