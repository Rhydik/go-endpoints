package api

import (
	"fmt"
	"main/internal/ports"
	"net/http"
)

type Adapter struct {
	mutexMap ports.MutexMapPort
}

func (httpa Adapter) SetMutexMap(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	httpa.mutexMap.Set(key, value)
	fmt.Fprintf(w, "Key: %s Value: %s successfully stored.", key, value)
}

func (httpa Adapter) GetMutexMap(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, ok := httpa.mutexMap.Get(key)

	if !ok {
		fmt.Fprintf(w, "Key not found")
		return
	}

	fmt.Fprintf(w, key, value)
}

func setupRoutes() {
	http.HandleFunc("/get", GetMutexMap)
	http.HandleFunc("/set", SetMutexMap)
}

func startHttp() {
	fmt.Println("HTTP Server Started on port 4000")
	setupRoutes()
	http.ListenAndServe(":4000", nil)
}
