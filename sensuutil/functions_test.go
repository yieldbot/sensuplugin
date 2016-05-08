package sensuutil

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
      "github.com/op/go-logging"

)

func TestPassword(t *testing.T) {

  Convey("When testing the redaction of a password", t, func() {
    p := "simple"
    Convey("If the password is 'simple'", func() {
      out := logging.Redact(string(p))
      Convey(" The output should be '******'", func() {
        So(out, ShouldEqual, "******")
        So(out, ShouldNotEqual, nil)
        So(out, ShouldNotEqual, "simple")
      })
    })
  } )
}
