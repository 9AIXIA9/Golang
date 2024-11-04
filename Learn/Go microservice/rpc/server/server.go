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

// 可以注册不同类型的多个对象（服务——结构体）但是不能定义同一类型 多个对象
//这里的类型指的是笼统的结构体，对象是实例

// ServiceA 自定义结构类型
type ServiceA struct {
}

// Add 为服务A增加一个可导出Add方法
func (s *ServiceA) Add(arg *Args, reply *int) (err error) {
	*reply = arg.X + arg.Y
	return nil
}

func main() {
	service := new(ServiceA)
	//注册RPC服务
	if err := rpc.Register(service); err != nil {
		log.Fatal("rpc register service failed,err:", err)
	}
	//rpc.HandleHTTP() //基于HTTP协议
	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen failed,err:", err)
	}
	//err = http.Serve(l, nil)
	//if err != nil {
	//	log.Fatal("serve failed,err:", err)
	//}
	//基于tcp协议处理RPC
	for {
		conn, _ := l.Accept()
		//rpc.ServeConn(conn)
		//替换成json协议
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
