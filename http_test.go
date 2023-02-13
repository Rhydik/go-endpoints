package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkSetEndpoint(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(httpSet))
	defer server.Close()

	for n := 0; n < b.N; n++ {
		res, _ := http.Get(server.URL + "/set?somekey=somevalue")
		res.Body.Close()
	}
}