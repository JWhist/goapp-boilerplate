package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JWhist/goapp-boilerplate/api"
	"github.com/JWhist/goapp-boilerplate/config"
	"github.com/flowchartsman/swaggerui"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewServer(c config.Config) *http.Server {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

	r := http.NewServeMux()

	r.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(api.Spec)))

	r.Handle("/metrics", promhttp.Handler())

	r.Handle("/", http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!!!\n"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})))

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
