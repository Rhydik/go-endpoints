package main

import (
	"einride_test/internal"
	"einride_test/internal/http"
	"einride_test/internal/memcached"
	"einride_test/internal/proto"
	"time"
)

func main() {
	mutexMap := internal.NewMutexMap()
	httpa := http.NewApplication(mutexMap)
	grpc := proto.MakegRPCServer(mutexMap)
	memcached := memcached.NewApplication(mutexMap)
	go httpa.StartApp()
	go grpc.StartgRPCServer()
	memcached.Listen()

	for {
		time.Sleep(1 * time.Second)
	}
}
