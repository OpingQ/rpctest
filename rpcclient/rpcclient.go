package rpcclient

import (
	"fmt"
	"log"
	"net/rpc"
	"rpctest/rpcstruct"
)

func getClient() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, err
}

func Arith_Multiply(args *rpcstruct.Args) *int {
	client, err := getClient()
	if err != nil {
		return nil
	}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("return nil", err)
		return nil
	}
	return &reply
}
