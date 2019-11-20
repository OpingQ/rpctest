package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"rpctest/grpccontnet/pb"

	"google.golang.org/grpc"
)

func MathMutli(request *pb.MathArgRequest) (*pb.MathMutliResponse, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil, errors.New(err.Error())
	}
	defer conn.Close()
	c := pb.NewMathServiceClient(conn)
	fmt.Println("")
	r, err := c.MathMutli(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New(err.Error())
	}
	return r, nil
}
