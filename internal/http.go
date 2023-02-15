package internal

import (
	"fmt"
	"net/http"
)

type MutexMapHTTPAdapter struct {
	mutexMap MutexMapPort
}

func (httpa *MutexMapHTTPAdapter) SetMutexMap(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.URL.Query() {
		httpa.mutexMap.SetMutexMap(key, value)
		fmt.Fprintf(w, "Key: %s Value: %s successfully stored.", key, value)
	}
}

func (httpa *MutexMapHTTPAdapter) GetMutexMap(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, ok := httpa.mutexMap.GetMutexMap(key)

	if !ok {
		fmt.Fprintf(w, "Key not found")
		return
	}

	fmt.Fprintf(w, "Key: %s Value: %s.", key, value)
}

func (httpa *MutexMapHTTPAdapter) setupRoutes() {
	http.HandleFunc("/get", httpa.GetMutexMap)
	http.HandleFunc("/set", httpa.SetMutexMap)
}

func NewApplication(httpa MutexMapHTTPAdapter) {
	fmt.Println("HTTP Server Started on port 4000")
	httpa.setupRoutes()
	http.ListenAndServe(":4000", nil)
}
