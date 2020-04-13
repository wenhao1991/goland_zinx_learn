package main

import (
	"fmt"
	"net"
	"time"
)

/*
	模拟客户端
 */
func main()  {
	fmt.Println("Client Start...")

	time.Sleep(1 * time.Second)
	// 1 直接连接远程服务器，得到conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8843")
	if err != nil{
		fmt.Println("Client start error, exit")
		return
	}
	for{
		//2 连接调用Write方法 写数据
		_, err :=conn.Write([]byte("Hello Znix"))
		if err != nil{
			fmt.Println("write error, err", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil{
			fmt.Println("read buff error, err", err)
			return
		}
		fmt.Printf("server call back:%s, cnt = %d\n", buf[:cnt], cnt)

		//cpu 阻塞
		time.Sleep(1*time.Second)
	}

}