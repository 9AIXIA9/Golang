package main

import (
	"add_client/pb"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	var x = flag.Int64("x", 10, "-x tell me x = ?")
	var y = flag.Int64("y", 10, "-y tell me y = ?")
	flag.Parse()
	log.Printf("x,y:%d %d", *x, *y)
	//连接server
	conn, err := grpc.Dial("127.0.0.1:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("grpc dial failed,err:%v", err)
		return
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal("conn close failed,err:", err)
		}
	}() //建立连接就要关闭
	//创建客户端
	c := pb.NewCounterClient(conn)
	//调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Add(ctx, &pb.AddRequest{
		X: *x,
		Y: *y,
	})
	if err != nil {
		log.Printf("add failed,err:%v", err)
		return
	}
	//处理RPC响应
	log.Printf("RPC response:%v", resp.Reply)
}
