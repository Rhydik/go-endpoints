package main

import (
	"einride_test/internal"
)

func main() {
	mutexMap := internal.NewMutexMap()

	httpa := internal.NewApplication(mutexMap)
	go httpa.StartApp()
	grpc := makeServer(mutexMap)
	grpc.startgRPCServer()
	//gRPC server
	//memcached whatever
}
