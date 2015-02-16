package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ip", getIPHandler)
	http.HandleFunc("/user-agent", getUserAgentHandler)
	http.HandleFunc("/headers", getHeadersHandler)
	http.HandleFunc("/get", getHandler)

	http.HandleFunc("/404", Handler404)

	http.ListenAndServe(":5000", nil)
}
