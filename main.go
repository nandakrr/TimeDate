// Golang program to get the current date and time

package main

import (
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now())
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
