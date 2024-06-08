//go:generate oapi-codegen --config=/Users/jordanwhistler/jw-goapp/api/config.yml /Users/jordanwhistler/jw-goapp/api/api.yml
package api

import (
	"encoding/json"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /ping)
func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
