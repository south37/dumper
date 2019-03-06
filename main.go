package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	defaultPort = "8080"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.Method, r.RequestURI, r.Proto)
	log.Printf("Host: %s", r.Host)
	for k, v := range r.Header {
		log.Printf("%s: %s", k, strings.Join(v, ","))
	}
	fmt.Fprintf(w, "Hello, dumper!")
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return p
	}
	return defaultPort
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/ping", handlePing)

	p := port()
	fmt.Printf("=> Server starting on :%s\n", p)
	if err := http.ListenAndServe(":"+p, nil); err != nil {
		panic(err)
	}
}
