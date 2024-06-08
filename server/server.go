package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JWhist/jw-goapp/api"
	"github.com/JWhist/jw-goapp/config"
)

func NewServer(c config.Config) *http.Server {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)
	return &http.Server{
		Addr:                         "0.0.0.0:" + fmt.Sprint(c.Port),
		DisableGeneralOptionsHandler: true,
		ReadTimeout:                  10 * time.Second,
		ReadHeaderTimeout:            5 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  300 * time.Second,
		Handler:                      h,
	}
}
