package eveapi

import (
	"net/http"
	"net/http/httptest"
)

const (
	testKey   = "1111111"
	testVCode = "ValidKeyValidKeyValidKeyValidKeyValidKeyValidKeyValidKeyValidKey"
)

// Returns a new server that serves the content of the file with the given name for requests to the given path.
func newServerFromFile(filename string, urlPath string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != urlPath {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, filename)
	}))
}
