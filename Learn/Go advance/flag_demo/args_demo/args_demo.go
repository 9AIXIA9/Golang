package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args是一个[]string
	//他为命令行参数
	fmt.Println(os.Args)
	if len(os.Args) > 0 {
		for i, arg := range os.Args {
			fmt.Printf("args[%d]:%v\n", i, arg)
		}
	}

}
