package main

import (
	"bufio"
	"code.xxx.com/backend/hello_client/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

//grpc client
//调用server的SayHello方法

func main() {
	//var name = flag.String("name", "xia", "-name tell me who are you?")
	//flag.Parse() //解析命令行参数
	//ServerStreamDemo(name)
	//ClientStreamDemo()
	BidiDemo()
}

// NormalDemo 普通调用RPC方法
func NormalDemo(name *string) {
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
	c := pb.NewGreeterClient(conn) //使用生成好的函数
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Printf("say hello failed,err:%v", err)
		return
	}
	//处理RPC响应
	log.Printf("RPC response:%v", resp.GetReply())
	//调用服务端流式RPC
}

func ServerStreamDemo(name *string) {
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
	c := pb.NewGreeterClient(conn) //使用生成好的函数
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Printf("get lots of replaces failed,err:%v", err)
		return
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream recv failed,err:%v", err)
			return
		}
		log.Printf("recv:%v\n", res.GetReply())
	}
}

func ClientStreamDemo() {
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
	c := pb.NewGreeterClient(conn) //使用生成好的函数
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//客户端要流式发送的数据
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Printf("lots of greeting failed,err:%v", err)
		return
	}
	name := []string{"张三", "李四", "王五"}
	for _, s := range name {
		err := stream.Send(&pb.HelloRequest{Name: s})
		if err != nil {
			log.Printf("stream send failed,err:%v", err)
			break
		}
	}
	//发送结束 关闭send
	res, err := stream.CloseAndRecv()
	if err != nil && err != io.EOF {
		log.Printf("stream close and recv failed,err:%v", err)
		return
	}
	log.Printf("server response:%v\n", res.GetReply())
}

func BidiDemo() {
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
	c := pb.NewGreeterClient(conn) //使用生成好的函数
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	//双流式模式
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Printf("bidi hello failed,err:%v", err)
		return
	}
	waitC := make(chan struct{})
	go func() {
		for {
			//接受服务端返回响应
			in, err := stream.Recv()
			if err == io.EOF {
				//接收完毕
				close(waitC)
				return
			}
			if err != nil {
				log.Fatalf("c.bidi hello failed,err:%v", err)
			}
			fmt.Printf("AI:%s\n", in.GetReply())
		}
	}()
	//从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) //从标注输入中生成读对象
	for {
		cmd, _ := reader.ReadString('\n') //读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		//将获取的数据发送到服务端
		err := stream.Send(&pb.HelloRequest{Name: cmd})
		if err != nil {
			log.Fatalf("stream send failed,err:%v", err)
		}
	}
	err = stream.CloseSend()
	if err != nil {
		log.Printf("stream close failed,err:%v", err)
	}
	<-waitC
}
