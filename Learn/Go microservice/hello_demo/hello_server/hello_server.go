package main

import (
	"context"
	"google.golang.org/grpc"
	"hello_server/pb"
	"log"
	"net"
)

//grpc server

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "Hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	//	启动服务
	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("net listen failed,err:", err)
	}
	s := grpc.NewServer() //创建grpc
	//注册服务
	pb.RegisterGreeterServer(s, &server{})
	//启动服务
	if err = s.Serve(l); err != nil {
		log.Fatal("serve failed,err:", err)
	}
}
