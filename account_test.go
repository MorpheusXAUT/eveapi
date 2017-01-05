package eveapi

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAPIKeyInfo(t *testing.T) {
	testServer := newServerFromFile("fixtures/account_APIKeyInfo.xml.aspx", "/account/APIKeyInfo.xml.aspx")
	defer testServer.Close()
	Convey("Setting up test server and parameters", t, func() {

		key := Key{ID: testKey, VCode: testVCode}
		api := &API{
			Server:    testServer.URL,
			APIKey:    key,
			UserAgent: "Hello",
			Debug:     false,
		}
		res, err := api.AccountAPIKeyInfo()
		Convey("Should not have errors", func() {
			So(err, ShouldBeNil)
		})
		Convey("Should be able to read the version", func() {
			So(res.Version, ShouldEqual, 2)
		})
		Convey("Test that some data loaded", func() {
			So(res.Key.Rows[0].ID, ShouldEqual, 1655827332)
			So(res.Key.Rows[0].Name, ShouldEqual, "Hel O'Ween")
			So(res.Key.Rows[0].CorporationID, ShouldEqual, 1226284052)
			So(res.Key.Rows[0].CorporationName, ShouldEqual, "Men On A Mission")
			So(res.Key.Rows[0].AllianceID, ShouldEqual, 0)
			So(res.Key.Rows[0].AllianceName, ShouldEqual, "")
			So(res.Key.Rows[0].FactionID, ShouldEqual, 0)
			So(res.Key.Rows[0].FactionName, ShouldEqual, "")
		})
	})

}

func TestAccountStatus(t *testing.T) {
	testServer := newServerFromFile("fixtures/account_AccountStatus.xml.aspx", "/account/AccountStatus.xml.aspx")
	defer testServer.Close()
	Convey("Setting up test server and parameters", t, func() {

		key := Key{ID: testKey, VCode: testVCode}
		api := &API{
			Server:    testServer.URL,
			APIKey:    key,
			UserAgent: "Hello",
			Debug:     false,
		}
		res, err := api.AccountStatus()
		Convey("Should not have errors", func() {
			So(err, ShouldBeNil)
		})
		Convey("Should be able to read the version", func() {
			So(res.Version, ShouldEqual, 2)
		})
		Convey("Test that some data loaded", func() {
			So(res.PaidUntil, ShouldNotBeEmpty)
			So(res.CreateDate, ShouldNotBeEmpty)
			So(res.LogonCount, ShouldEqual, 3177)
			So(res.LogonMinutes, ShouldEqual, 235243)

		})
	})

}
func TestAccountStatusFail(t *testing.T) {
	testServer := newServerFromFile("fixtures/fail", "/fail")
	defer testServer.Close()
	Convey("Setting up test server and parameters", t, func() {

		key := Key{ID: testKey, VCode: testVCode}
		api := &API{
			Server:    testServer.URL,
			APIKey:    key,
			UserAgent: "Hello",
			Debug:     false,
		}
		_, err := api.AccountStatus()
		Convey("Should not have errors", func() {
			So(err, ShouldNotBeNil)
		})
	})

}

func TestAccountGetChars(t *testing.T) {
	testServer := newServerFromFile("fixtures/account_APIKeyInfo.xml.aspx", "/account/APIKeyInfo.xml.aspx")
	defer testServer.Close()
	Convey("Setting up test server and parameters", t, func() {

		key := Key{ID: testKey, VCode: testVCode}
		api := &API{
			Server:    testServer.URL,
			APIKey:    key,
			UserAgent: "Hello",
			Debug:     false,
		}
		res, err := api.AccountGetChars()
		Convey("Should not have errors", func() {
			So(err, ShouldBeNil)
		})
		Convey("Test that some data loaded", func() {
			So(res[0], ShouldEqual, "Hel O'Ween")
		})
	})

}

func TestAccountGetCharsFail(t *testing.T) {
	testServer := newServerFromFile("fixtures/fail", "/fail")
	defer testServer.Close()
	Convey("Setting up test server and parameters", t, func() {

		key := Key{ID: testKey, VCode: testVCode}
		api := &API{
			Server:    testServer.URL,
			APIKey:    key,
			UserAgent: "Hello",
			Debug:     false,
		}
		res, err := api.AccountGetChars()
		Convey("Should not have errors", func() {
			So(err, ShouldNotBeNil)
		})
		Convey("Test that some data loaded", func() {
			So(res, ShouldBeNil)
		})
	})

}
