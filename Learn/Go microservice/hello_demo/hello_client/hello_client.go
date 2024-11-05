package main

import (
	"code.xxx.com/backend/hello_client/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

//grpc client
//调用server的SayHello方法

func main() {
	var name = flag.String("name", "xia", "-name tell me who are you?")
	flag.Parse() //解析命令行参数
	//连接server
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc dial failed")
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal("conn close failed,err:", err)
		}
	}()
	//创建客户端
	c := proto.NewGreeterClient(conn) //使用生成好的函数
	//调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Printf("say hello failed,err:%v", err)
		return
	}
	//处理RPC响应
	log.Printf("RPC response:%v", resp.GetReply())
}
