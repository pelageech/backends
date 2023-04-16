package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("cache-control", "max-age=10")
	fmt.Println(r)
	time.Sleep(5 * time.Second)
	if _, err := fmt.Fprintln(w, time.Now()); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
func main() {

	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(":3032", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
