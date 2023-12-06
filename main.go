package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	_ = r.Run(":8080")
}
