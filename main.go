package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<h1 style="color:blue;">V0.0.1</h1>`)
	})
	log.Printf("Serving on HTTP port 80")
	http.ListenAndServe(":80", nil)
}
