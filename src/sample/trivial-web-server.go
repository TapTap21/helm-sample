package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am a GO application running inside Docker. - 0.14")

}

func main() {
	fmt.Println("Basic web server is starting on port 80...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}
