package main

import (
	"fmt"
	"os"
	"rpctest/grpccontnet/client"
	"rpctest/grpccontnet/pb"
	"rpctest/grpccontnet/server"
	"rpctest/rpcclient"
	"rpctest/rpcserver"
	"rpctest/rpcstruct"
)

func main() {
	fmt.Println(os.Args)
	switch os.Args[1] {
	case "server":
		rpcserver.StartRpc()
	case "client":
		args := &rpcstruct.Args{A: 5, B: 8}
		reply := rpcclient.Arith_Multiply(args)
		fmt.Printf("Arith: %d*%d=%d", args.A, args.B, *reply)
	case "grpc_server":
		server.StartGRPCServer()
	case "grpc_client":
		res, _ := client.MathMutli(&pb.MathArgRequest{A: 8, B: 7})
		fmt.Println(res)
	}
}

func testMain() {
	fmt.Println("測試衝突1")
}
