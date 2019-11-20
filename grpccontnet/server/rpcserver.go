package server

import (
	"log"
	"net"
	"rpctest/grpccontnet/pb"
	"rpctest/grpccontnet/server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer() {

	// 新增Server
	s := grpc.NewServer()

	// 註冊服務
	pb.RegisterMathServiceServer(s, &service.Server{})

	// 要監聽的port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	// 啟動Server
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
