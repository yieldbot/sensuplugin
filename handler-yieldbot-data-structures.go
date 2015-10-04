// Library for all data structures used in Yieldbot alert handlers and dashboard generators
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package dracky

// Data structure for holding Sensu generated check results.
type Sensu_Event struct {
	Action      string
	Occurrences int
	Client      struct {
		Name          string
		Address       string
		Subscriptions []string
		Timestamp     int64
	}
	Check struct {
		Source      string
		Name        string
		Issued      int64
		Subscribers []string
		Interval    int
		Command     string
		Output      string
		Status      int
		Handler     string
		History     []string
	}
}

// Data structure for holding environment variables provided by Oahi dropped via Chef.
type Env_Details struct {
	Sensu struct {
		Environment string `json:"environment"`
		FQDN        string `json:"fqdn"`
		Hostname    string `json:"hostname"`
	}
}

// Data structure for holding generic user data that is entered via an input file declared on the commndline.
type User_Event struct {
	Product   string
	Timestamp int64
	Data      string
}
