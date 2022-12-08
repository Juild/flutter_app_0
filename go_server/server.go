package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server Reached")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Reached")
}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/getUser", getUser)

	http.ListenAndServe(":80", nil)
}
