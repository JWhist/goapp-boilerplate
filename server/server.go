package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jwhist/jw-goapp/config"
)

func NewServer(c config.Config) *http.Server {
	return &http.Server{
		Addr:                         "0.0.0.0:" + fmt.Sprint(c.Port),
		DisableGeneralOptionsHandler: true,
		ReadTimeout:                  10 * time.Second,
		ReadHeaderTimeout:            5 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  300 * time.Second,
		Handler:                      nil,
	}
}
