package ping

import (
	"fmt"
	"net/http"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("ping: %v", "pong")))
		w.WriteHeader(http.StatusOK)
		return
	}
}
