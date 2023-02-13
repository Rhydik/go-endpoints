package main

import (
	"fmt"
	"net/http"
)

func httpSet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	mutexMap.set(key, value)
	fmt.Fprintf(w, "Key: %s Value: %s successfully stored.", key, value)
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, ok := mutexMap.get(key)

	if !ok {
		fmt.Fprintf(w, "Key not found")
		return
	}

	fmt.Fprintf(w, key, value)
}

func setupRoutes() {
	http.HandleFunc("/get", httpGet)
	http.HandleFunc("/set", httpSet)
}

func startHttp() {
	fmt.Println("HTTP Server Started on port 4000")
	setupRoutes()
	http.ListenAndServe(":4000", nil)
}
