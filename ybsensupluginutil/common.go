// Library for all global variables used by the Yieldbot
// Infrastructure teams in sensu
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package ybsensupluginutil

// Debug  Do we print debug statements or not? This is set in each binary but is placed here
// to avoid the use of global variables
var Debug bool

// MonitoringErrorCodes provides a standard set of error codes to use.
// Please use the below codes instead of random non-zero so that monitoring can
// utilize existing maps for alerting and help avoid unnecessary noise.
var MonitoringErrorCodes = map[string]int{
	"GeneralGolangError": 129,
	"ConfigError":        127,
	"PermissionError":    126,
	"RuntimeError":       42,
	"OK":                 1,
	"WARNING":            2,
	"CRITICAL":           3,
	"UNKNOWN":            4,
}
