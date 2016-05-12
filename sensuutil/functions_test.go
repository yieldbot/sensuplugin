package sensuutil

import (
	"testing"

	"github.com/op/go-logging"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPassword(t *testing.T) {

	Convey("When testing the redaction of a password", t, func() {
		p := "simple"
		Convey("If the password is 'simple'", func() {
			out := logging.Redact(string(p))
			Convey("The output should be '******'", func() {
				So(out, ShouldEqual, "******")
				So(out, ShouldNotEqual, nil)
				So(out, ShouldNotEqual, "simple")
			})
		})
	})
}

// func TestExit(t *testing.T) {
// 	// var exitCode string
// 	// var exitOutput string
// 	Convey("When testing a exit code with no output", t, func() {
// 		// exitOutput = ""
// 		Convey("If the exit text is 'critical'", func() {
// 			Exit("critical")
// 			Convey("The exit code shoult be 3", func() {
// 				So(Exit, ShouldEqual, 3)
// 			})
// 		})
// 	})
// }
