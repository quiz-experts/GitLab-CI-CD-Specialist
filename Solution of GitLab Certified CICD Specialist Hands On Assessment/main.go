package main

import (
	"fmt"
	"net/http"
)

func helloworld() string {
	return "Hello World!!"
}

func healthcheck() string {
	return "Health OK!"
}

func livenesscheck() string {
	return "I am alive!!!"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, helloworld())
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, healthcheck())
	})

	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, livenesscheck())
	})

	http.ListenAndServe(":8080", nil)
}