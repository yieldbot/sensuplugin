package sensuutil

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  "errors"
  "fmt"

)

func TestEhndlr(t *testing.T) {

  Convey("When testing the printing of an error", t, func() {
    e := errors.New("no joy")
    Convey("If there is an error message", func() {
      EHndlr(e)

      Convey(" The output should be error", func() {
        So(EHndlr(e), ShouldEqual, "no joy")
        So(e, ShouldNotEqual, nil)
        So(e, ShouldNotEqual, "")
      })
    })


  } )



}
