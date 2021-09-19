package proxy

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCache struct{}

func (c *testCache) Get(key KeyCache) (*CacheData, bool) {
	return nil, false
}

func (c *testCache) Put(key KeyCache, data *CacheData) {

}

func TestProxy(t *testing.T) {
	tests := []struct {
		proxy      *Proxy
		statusCode int
		body       string
	}{
		{New(Params{"https", "github.com", 443, false, []string{"GET", "HEAD"}}, &testCache{}), 200, "github"},
		{New(Params{"https", "github.com", 443, true, []string{"GET", "HEAD"}}, &testCache{}), 200, "github"},
		{New(Params{"https", "google.com", 443, false, []string{"GET", "HEAD"}}, &testCache{}), 200, "google"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		test.proxy.Handle(rr, req)
		if status := rr.Code; status != test.statusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if !strings.Contains(rr.Body.String(), test.body) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.body)
		}
	}

}
