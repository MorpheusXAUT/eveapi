package eveapi

import (
	"encoding/xml"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	_ "runtime"
	"testing"
)

func TestCorpWalletJournal(t *testing.T) {
	testServer := newServerFromFile("fixtures/corp_WalletJournal.xml.aspx", "/corp/WalletJournal.xml.aspx")
	defer testServer.Close()
	Convey("Setting up test server and parameters", t, func() {

		key := Key{ID: testKey, VCode: testVCode}
		api := &API{
			Server:    testServer.URL,
			APIKey:    key,
			UserAgent: "Hello",
			Debug:     false,
		}

		//runtime.Breakpoint()
		res, err := api.CorpWalletJournal(1000, 0, 10)
		Convey("Should not have errors", func() {
			So(err, ShouldBeNil)
		})

		Convey("Should be able to read the version", func() {
			So(res.Version, ShouldEqual, 2)
		})

		Convey("Test that some data loaded", func() {
			tx := res.Transactions[0]
			So(tx.RefID, ShouldBeGreaterThan, 0)
		})
	})

}

func TestCorpWalletJournalDecode(t *testing.T) {
	Convey("Decoding an API response", t, func() {
		xmlFile, err := os.Open("fixtures/corp_WalletJournal.xml.aspx")
		Convey("Should read fixtures", func() {
			So(err, ShouldBeNil)
		})

		output := CorpWalletJournalResult{}
		xml.NewDecoder(xmlFile).Decode(&output)

		Convey("Should decode responses", func() {
			So(output.Version, ShouldEqual, 2)
		})

	})

}
