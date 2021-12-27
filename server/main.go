package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world from dew-server")
}
