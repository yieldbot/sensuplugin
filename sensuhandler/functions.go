// Library for all handler functions used by the Yieldbot Infrastructure
// teams in sensu.
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

// Package sensuhandler implements common data structures and functions for Yieldbot monitoring alerts and dashboards
package sensuhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/yieldbot/sensuplugin/sensuutil"
)

// AcquireUchiwa returns an uchiwa url for the node alerting
// func (e SensuEvent) AcquireMonitoredInstance() string {
func (e EnvDetails) AcquireUchiwa(h string) string {
	var tags string
	var dc string

	tags = e.Sensu.Consul.Tags
	dc = e.Sensu.Consul.Datacenter
	fmt.Println(e.Sensu.Consul.Tags)
	fmt.Println(e.Sensu.Consul.Datacenter)

	fmt.Println(tags)
	fmt.Println(e)
	fmt.Println(dc)
	fmt.Println(h)

	url := "https://" + tags + ".uchiwa.service" + "." + dc + ".consul/#/client/" + dc + "/" + h

	return url
}

// CleanOutput will shorten the output to a more manageable length
func CleanOutput(output string) string {
	return strings.Split(output, ":")[0]
}

// EventName generates a simple string for use by elasticsearch and internal logging of all monitoring alerts.
func EventName(client string, check string) string {
	return client + "_" + check
}

// AcquireMonitoredInstance sets the correct device that is being monitored. In the case of snmp trap collection, containers,
// or appliance monitoring the device running the sensu-client may not be the device actually being monitored.
func (e SensuEvent) AcquireMonitoredInstance() string {
	if e.Check.Source != "" {
		return e.Check.Source
	}
	return e.Client.Name
}

// SetColor is used to set the correct notification color for a given status. By setting it in a single place for all alerts
// we ensure a measure of cohesiveness across various notification channels.
func SetColor(status int) string {
	switch status {
	case 0:
		return NotificationColor["green"]
	case 1:
		return NotificationColor["yellow"]
	case 2:
		return NotificationColor["red"]
	case 3:
		return NotificationColor["orange"]
	default:
		return NotificationColor["orange"]
	}
}

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
	eCode := "UNDEFINED_STATUS"

	for k, v := range sensuutil.MonitoringErrorCodes {
		if status == v {
			eCode = k
		}
	}
	return eCode
}

// CreateCheckName creates a monitor name that is easily searchable in ES using different
// levels of granularity.
func CreateCheckName(check string) string {
	return strings.Replace(check, "-", ".", -1)
}

// DefineCheckStateDuration calculates how long a monitor has been in a given state.
func DefineCheckStateDuration() int {
	return 0
}

// SetSensuEnv reads in the environment details provided by Oahi and drops them into a statically defined struct.
func (e EnvDetails) SetSensuEnv() *EnvDetails {
	envFile, err := ioutil.ReadFile(sensuutil.EnvironmentFile)
	if err != nil {
		panic(err)
	}

	// var envDetails EnvDetails
	err = json.Unmarshal(envFile, &e)
	if err != nil {
		panic(err)
	}
	return &e
}

// AcquireSensuEvent reads in the check result generated by Sensu via stdin and drop it into a statically defined struct.
func (e SensuEvent) AcquireSensuEvent() *SensuEvent {
	results, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(results, &e)
	if err != nil {
		panic(err)
	}
	return &e
}
