package main

import (
	"context"

	"google.golang.org/grpc"
)

func client_main() {
	conn, err := grpc.Dial("localhost:4001", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := NewMutexMapClient(conn)
	res, err := c.GetMutexMap(context.Background(), &KeyRequest{Key: "somekey"})
	if err != nil {
		panic(err)
	}
	println(res.Value)
}
