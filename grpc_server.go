package main

import (
	"context"
	"net"

	"einride_test/internal"

	"google.golang.org/grpc"
)

type server struct {
	mutexMap internal.MutexMapPort
}

func (s *server) GetMutexMap(ctx context.Context, in *KeyRequest) (*ValueReply, error) {
	vr := &ValueReply{}

	val, ok := s.mutexMap.GetMutexMap(in.Key)
	if !ok {
		vr.Value = "Key not found"
		return vr, nil
	}
	vr.Value = val[0]

	return vr, nil
}

func (s *server) SetMutexMap(ctx context.Context, in *KeyRequest) (*Empty, error) {
	s.mutexMap.SetMutexMap(in.Key, []string{in.Value})
	return &Empty{}, nil
}

func (s *server) mustEmbedUnimplementedMutexMapServer() {}

func makeServer(mutexMap internal.MutexMapPort) *server {
	return &server{mutexMap: mutexMap}
}

func (s *server) startgRPCServer() {
	lis, err := net.Listen("tcp", ":4001")
	if err != nil {
		panic(err)
	}

	grpcS := grpc.NewServer()
	RegisterMutexMapServer(grpcS, &server{s.mutexMap})
	err = grpcS.Serve(lis)
	if err != nil {
		panic(err)
	}
}
