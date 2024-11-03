package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP client demo
func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("dial failed,err:", err)
		return
	}
	defer c.Close()
	for {
		input := bufio.NewReader(os.Stdin)
		s, _ := input.ReadString('\n')
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Println("write failed,err:", err)
			return
		}
		//接收数据
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read from udp failed,err:", err)
			return
		}
		fmt.Printf("read from %v,mag:%v\n", addr, string(buf[:n]))
	}
}
