package rpcserver

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpctest/rpcstruct"
	"time"
)

func StartRpc() {
	for _, v := range rpcStructList() {
		rpc.Register(v)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error:", e)
	}

	go http.Serve(l, nil)
	for {
		time.Sleep(2 * time.Second)
	}
}

func rpcStructList() map[int]interface{} {
	return map[int]interface{}{
		1: new(rpcstruct.Arith),
		2: new(rpcstruct.Arith2),
	}
}
