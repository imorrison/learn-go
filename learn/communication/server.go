package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"][0]
	fmt.Fprintf(w, "Ping Pong Player is %s \n", name)
}

func main() {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}
