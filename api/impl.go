//go:generate oapi-codegen --config=/Users/jordanwhistler/jw-goapp/api/config.yml /Users/jordanwhistler/jw-goapp/api/api.yml
package api

import (
	_ "embed"
	"encoding/json"
	"net/http"
)

type Server struct{}

//go:embed api.yml
var Spec []byte

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
