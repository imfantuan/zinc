package core

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewIndex(t *testing.T) {
	Convey("test new index storage dick", t, func() {
		indexName := "create.new.index"
		index, err := NewIndex(indexName, "disk")
		So(err, ShouldBeNil)
		So(index.Name, ShouldEqual, indexName)
	})
	Convey("test new index storage s3", t, func() {
		// TODO: support
	})
}
