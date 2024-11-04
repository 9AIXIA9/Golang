package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

func main() {
	//建立HTTP连接
	//client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	//基于TCP建立RPC连接
	//client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	//基于JSON协议
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	if err != nil {
		log.Fatal(err)
	}
	//同步调用
	args := &Args{
		X: 10,
		Y: 20,
	}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Service.Add:%d+%d = %d", args.X, args.Y, reply)

	//异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done //调用结果通知
	log.Print(replyCall.Error)
	log.Print(reply2)
}
