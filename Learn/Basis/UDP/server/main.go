package main

import (
	"fmt"
	"net"
)

//UDP server demo

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return

	}
	defer listen.Close()
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read from udp failed,err:", err)
			return
		}
		fmt.Println("接收到的数据", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to %v failed,err:%v\n", addr, err)
			return
		}
	}
}
