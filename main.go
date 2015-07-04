package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Mux struct {
	http.Handler
}

func (mux Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	domainParts := strings.Split(r.Host, ".")
	fmt.Println(domainParts)

	mux.ServeHTTP(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello, %q", r.URL.Path[1:])
}

func main() {

	mux := http.NewServeMux()
	//mux.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":80", mux)

	/*
		http.Handle("/", http.FileServer(http.Dir("static")))
		http.ListenAndServe(":555", nil)
	*/
}
