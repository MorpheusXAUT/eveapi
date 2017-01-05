package eveapi

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleClient(t *testing.T) {
	Convey("All parameters should pass", t, func() {
		testKey := "1118648"
		testVCode := "ValidKey77kMPi96cpK5K9vp7GAdQh9TzmDyfooKRbSFhplfNAceQSQ3vWWMuSGh"
		key := Key{ID: testKey, VCode: testVCode}
		api := Simple(key)
		So(api.Server, ShouldEqual, Tranquility)
		So(api.APIKey.ID, ShouldEqual, testKey)
		So(api.APIKey.VCode, ShouldEqual, testVCode)
		So(api.UserAgent, ShouldEqual, "Go API Wrapper")
	})

}
