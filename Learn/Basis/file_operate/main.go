package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件操作
// 读取文件所有信息
func readAll() {
	//相对路径，相对可执行程序的目录下的xxx.txt
	//打开文件
	fileObj, err := os.Open("E:\\Go\\Learn\\file_operate\\xxx.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return

	}
	//关闭文件
	defer fileObj.Close()
	//读取文件
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		//读到文件末尾
		if err == io.EOF {
			//把当前读了多少个字节打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp))
	}
}

// read by bufio
func readBufio() {
	//相对路径，相对可执行程序的目录下的xxx.txt
	//打开文件
	fileObj, err := os.Open("E:\\Go\\Learn\\file_operate\\xxx.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return

	}
	//关闭文件
	defer fileObj.Close()
	//读取文件
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read file by bufio failed,err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil ---> readFile
func readByIoutil() {
	content, err := os.ReadFile("E:\\Go\\Learn\\file_operate\\xxx.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed,err:%v\n", err)
		return
	}
	fmt.Println(string(content))
}

// write
func write() {
	fileObj, err := os.OpenFile("E:\\Go\\Learn\\file_operate\\xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "我是小王子 "
	fileObj.Write([]byte(str))        //[]byte
	fileObj.WriteString("hello 小王子 ") //string
}

// write by bufio
func writeByBufio() {
	fileObj, err := os.OpenFile("E:\\Go\\Learn\\file_operate\\xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("小王子，你好呀") //将内容写入缓冲区
	writer.Flush()                //将缓冲区内容写入文件
}

// write by ioutil
func writeByIoutil() {
	str := "人生得意须尽欢"
	//ioutil.WriteFile ----> os.WriteFile
	err := os.WriteFile("E:\\Go\\Learn\\file_operate\\xxx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed,err:%v\n", err)
		return
	}

}
func main() {
	//readAll()
	//readBufio()
	//readByIoutil()
	//write()
	//writeByBufio()
	writeByIoutil()
}
