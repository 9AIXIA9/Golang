package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"hello_server/pb"
	"io"
	"log"
	"net"
	"strings"
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

func (s server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好 ",
		"hello ",
		"空你几哇 ",
		"阿尼啊谁有 ",
	}
	for _, word := range words {
		data := &pb.HelloResponse{Reply: word + in.GetName()}
		//使用send方法返回多个方法
		if err := stream.Send(data); err != nil {
			log.Printf("stream send failed,err：%v", err)
		}
	}
	return nil
}

func (s server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
	reply := "你好呀!"
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			//发送完统一回复
			if err := stream.SendAndClose(&pb.HelloResponse{Reply: reply}); err != nil {
				log.Printf("stream send and close failed,err:%v", err)
				return nil
			}
			break
		}
		if err != nil {
			log.Printf("stream recv failed,err:%v\n", err)
			break
		}
		reply += res.GetName() + " "
	}
	log.Printf("reply:%v", reply)
	return nil
}

func (s server) BidiHello(stream pb.Greeter_BidiHelloServer) error {
	for {
		//接受流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			//发送完毕
			return nil
		}
		if err != nil {
			//出错
			log.Printf("stream recv failed,err:%v", err)
			return nil
		}
		reply := logic(in.GetName()) //对收到的信息做处理
		//返回流式响应
		err = stream.Send(&pb.HelloResponse{Reply: reply})
		if err != nil {
			log.Printf("stream send failed,err:%v", err)
			return err
		}
	}
}

// 做一些处理
func logic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "?", "!")
	s = strings.ReplaceAll(s, "？", "！")
	return s
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
