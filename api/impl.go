//go:generate oapi-codegen --config=/Users/jordanwhistler/goapp-boilerplate/api/config.yml /Users/jordanwhistler/goapp-boilerplate/api/api.yml
package api

import (
	_ "embed"
	"encoding/json"
	"net/http"
)

// Server implements the API server interface.
type Server struct{}

// Spec contains the embedded OpenAPI specification.
//
//go:embed api.yml
var Spec []byte

// NewServer creates a new instance of the API server.
func NewServer() Server {
	return Server{}
}

// GetPing handles the GET /ping endpoint.
func (Server) GetPing(w http.ResponseWriter, _ *http.Request) {
	resp := Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
