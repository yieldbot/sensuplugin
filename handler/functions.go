// Library for all functions used by the Yieldbot Infrastructure teams
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

// Package dracky implements common data structures and functions for Yieldbot monitoring alerts and dashboards
package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// EventName generates a simple string for use by elasticsearch and internal logging of all monitoring alerts.
func EventName(client string, check string) string {
	return client + "_" + check
}

// AcquireMonitoredInstance sets the correct device that is being monitored. In the case of snmp trap collection, containers, or applicance
// monitoring the device running the sensu-client may not be the device actually being monitored.
func (e SensuEvent) AcquireMonitoredInstance() string {
	var monitoredInstance string
	if e.Check.Source != "" {
		monitoredInstance = e.Check.Source
	} else {
		monitoredInstance = e.Client.Name
	}
	return monitoredInstance
}

// func Set_time(t int) string {
//
// 	timeStamp := time.Unix(unixIntValue, 0)
// 	timestamp = time.Unix(timestamp, 0).Format(time.RFC822Z)
//
// }

// DefineSensuEnv sets the environment that the machine is running in based upon values
// dropped via Oahi during the Chef run.
func DefineSensuEnv(env string) string {
	switch env {
	case "prd":
		return "Prod "
	case "dev":
		return "Dev "
	case "stg":
		return "Stg "
	case "vagrant":
		return "Vagrant "
	default:
		return "Test "
	}
}

// DefineStatus converts the check result status from an integer to a string.
func DefineStatus(status int) string {
	switch status {
	case 0:
		return "OK"
	case 1:
		return "WARNING"
	case 2:
		return "CRITICAL"
	case 3:
		return "UNKNOWN"
	case 126:
		return "PERMISSION DENIED"
	case 127:
		return "CONFIG ERROR"
	default:
		return "UNKNOWN"
	}
}

// CreateCheckName creates a monitor name that is easliy searchable in ES using different
// levels of granularity.
func CreateCheckName(check string) string {
	fmtdCheck := strings.Replace(check, "-", ".", -1)
	return fmtdCheck
}

// DefineCheckStateDuration calculates how long a monitor has been in a given state.
func DefineCheckStateDuration() int {
	return 0
}

// SetSensuEnv reads in the environment details provided by Oahi and drop it into a staticly defined struct.
func SetSensuEnv() *EnvDetails {
	envFile, err := ioutil.ReadFile(EnvironmentFile)
	if err != nil {
		Check(err)
	}

	var envDetails EnvDetails
	err = json.Unmarshal(envFile, &envDetails)
	if err != nil {
		Check(err)
	}
	return &envDetails
}

// AcquireSensuEvent reads in the check result generated by Snesu via stdin and drop it into a staticaly defined struct.
func (e SensuEvent) AcquireSensuEvent() *SensuEvent {
	results, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		Check(err)
	}
	err = json.Unmarshal(results, &e)
	if err != nil {
		Check(err)
	}
	return &e
}

// Check is for generic error handling in all Yieldbot alert and dashboard packages.
func Check(e error) {
	if e != nil {
		fmt.Printf("%v", e)
		panic(e)
	}
}