package main

import (
	"add_server/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	*pb.UnimplementedCounterServer
}

func (s Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	reply := req.X + req.Y
	return &pb.AddResponse{Reply: reply}, nil
}

func main() {
	//启动监听服务
	l, err := net.Listen("tcp", "127.0.0.1:9092")
	if err != nil {
		log.Fatal("net listen failed,err:", err)
	}
	s := grpc.NewServer() //创建grpc
	//注册服务
	pb.RegisterCounterServer(s, &Server{})
	//启动服务
	err = s.Serve(l)
	if err != nil {
		log.Fatal("s serve failed,err:", err)
	}
}
