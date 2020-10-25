package ping

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	srv := httptest.NewServer(Ping())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/ping", srv.URL))
	if err != nil {
		t.Fatalf("couldn to send a get request to %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}
