package main

import (
	"fmt"
	"github.com/blankstars/learn_gee/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world: %v", req.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "HEADER[%q] = %q\n", k, v)
		}
	})

	r.Run(":8080")
}
