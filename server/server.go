// Package server provides HTTP server functionality for the application.
package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/JWhist/goapp-boilerplate/api"
	"github.com/JWhist/goapp-boilerplate/config"
	"github.com/JWhist/goapp-boilerplate/static"
	"github.com/flowchartsman/swaggerui"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewServer creates and configures a new HTTP server with the provided configuration.
func NewServer(c config.Config) *http.Server {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

	r := http.NewServeMux()

	r.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(api.Spec)))

	r.Handle("/metrics", promhttp.Handler())

	// Get the UI handler
	uiHandler := static.GetUIHandler()

	// Handle static assets (CSS, JS, etc.)
	r.Handle("/static/", uiHandler)

	// Custom handler for root and SPA routing
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		// Serve API routes (let the API handler deal with them later)
		if strings.HasPrefix(path, "/ping") || strings.HasPrefix(path, "/api/") {
			// Let it fall through to the API handler
			http.NotFound(w, req)
			return
		}

		// Serve static files and handle SPA routing
		uiHandler.ServeHTTP(w, req)
	})

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
