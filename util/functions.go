// Library for general purpose functions used by the Yieldbot Infrastructure
// teams in sensu.
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package lib

// Check is for generic error handling in all Yieldbot alert and dashboard packages.
func Check(e error) {
	if e != nil {
		fmt.Printf("%v", e)
		panic(e)
	}
}
