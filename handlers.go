package main

import (
	"encoding/json"
	"net"
	"net/http"
)

type JsonResponse map[string]interface{}

func (r JsonResponse) Serialize() []byte {
	serializedData, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return []byte("404 Not Found")
	}
	return serializedData
}

func indexHandler(rw http.ResponseWriter, request *http.Request) {}

func getIPHandler(rw http.ResponseWriter, request *http.Request) {
	ip, _, _ := net.SplitHostPort(request.RemoteAddr)
	response := JsonResponse{
		"origin": ip,
	}
	rw.Write(response.Serialize())
}

func getUserAgentHandler(rw http.ResponseWriter, request *http.Request) {
	userAgent := request.UserAgent()
	response := JsonResponse{
		"user-agent": userAgent,
	}
	rw.Write(response.Serialize())
}

func getHeadersHandler(rw http.ResponseWriter, request *http.Request) {
	requestHeaders := request.Header
	headers := make(map[string]string)
	for key := range requestHeaders {
		headers[key] = requestHeaders.Get(key)
	}
	data := map[string]map[string]string{
		"headers": headers,
	}
	serializedData, _ := json.MarshalIndent(data, "", "    ")
	rw.Write(serializedData)
}

func Handler404(rw http.ResponseWriter, request *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte("404 Not Found"))
}

func getHandler(rw http.ResponseWriter, request *http.Request) {
	data := JsonResponse{
		"origin": 23,
		"hello":  true,
		"world":  "yes",
	}
	serializedData, _ := json.MarshalIndent(data, "", "    ")
	rw.Write(serializedData)
}
