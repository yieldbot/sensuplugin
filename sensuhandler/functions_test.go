package cmd

import . "github.com/smartystreets/goconvey/convey"

//
// func AcquireUchiwa(h string, env interface{}) string {
// 	var tags string
// 	var dc string
//
// 	if e, ok := env.(EnvDetails); ok {
// 		tags = e.Sensu.Consul.Tags
// 		dc = e.Sensu.Consul.Datacenter
// 	}
//
// 	url := "https://" + tags + ".uchiwa.service" + "." + dc + ".consul/#/client/" + dc + "/" + h
//
// 	return url
// }
//
// func TestAcquireUchiwa(t *testing.T) {
// 	var tags string
// 	var dc string
//   var host string
//   var url string
//
//   sensuEvent := new(sensuhandler.SensuEvent)
//
//   type EnvDetails struct {
//   	Sensu struct {
//   		Environment string `json:"environment"`
//   		FQDN        string `json:"fqdn"`
//   		Hostname    string `json:"hostname"`
//   		Consul      struct {
//   			Tags       string `json:"tags"`
//   			Datacenter string `json:"datacenter"`
//   		}
//   	}
//   }
//
//
// 	Convey("When generating an Uchiwa string", t, func() {
//
// 		Convey("If consul tag is 'general' and the dc is 'us-atlant-1'", func() {
// 			tags = "general"
//       dc = "us-atlanta-1"
//       host = "dev-test"
//
//       url = AcquireUchiwa(host, env interface{})
//
//
//
// 			Convey("The alert status should be critical and the message should not be empty", func() {
// 				So(condition, ShouldEqual, "critical")
// 				So(msg, ShouldEqual, "Chrony is synced locally")
// 			})
//
// 			Convey("The alert status should not be ok and the message should be empty", func() {
// 				So(condition, ShouldNotEqual, "ok")
// 				So(msg, ShouldNotBeEmpty)
// 			})
// 		})
// 		Convey("If the RefID is a remote IP", func() {
// 			RefID = "8.8.8.8"
// 			condition, msg = checkLocalChrony(RefID)
//
// 			Convey("The alert status should be ok and the message should be empty", func() {
// 				So(condition, ShouldEqual, "ok")
// 				So(msg, ShouldBeEmpty)
// 			})
//
// 			Convey("The alert status should not be critical and the message should not be empty", func() {
// 				So(condition, ShouldNotEqual, "critical")
// 				So(msg, ShouldNotEqual, "Chrony is synced locally")
// 			})
// 		})
// 	})
// }
