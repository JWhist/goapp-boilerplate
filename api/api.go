package api

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	fmt.Printf("Request received [%v] %v %v %v\n", time.Now().Format(time.Stamp), r.Proto, r.Method, r.URL.Path)
}
