package service

import (
	"context"
	"rpctest/grpccontnet/pb"
)

type Server struct{}

func (s *Server) MathMutli(c context.Context, request *pb.MathArgRequest) (*pb.MathMutliResponse, error) {
	result := &pb.MathMutliResponse{}

	result.Ab1 = request.A * request.B * 1
	result.Ab2 = request.A * request.B * 2

	return result, nil
}
