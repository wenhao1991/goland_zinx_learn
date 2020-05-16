package main

import (
	"awesomeProject/Zinx/Znet"
	"fmt"
	"io"
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
	conn, err := net.Dial("tcp", "127.0.0.1:8433")
	if err != nil{
		fmt.Println("Client start error, exit")
		return
	}
	for{
		// send packed message, MsgId: 0
		dp := Znet.NewDataPack()
		binaryMsg, err :=dp.Pack(Znet.NewMessage(0, []byte("ZinxV0.5 client Test Message")))
		if err != nil{
			fmt.Println("Client pack error, exit")
			return
		}
		if _, err :=conn.Write(binaryMsg); err != nil{
			fmt.Println("Client write error, exit")
			return
		}

		// server should reply message to us

		// read head from stream, get Id and datalen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil{
			fmt.Println("read head error", err)
			break
		}

		// read
		msgHead, err :=dp.Unpack(binaryHead)
		if err != nil{
			fmt.Println("client unpack msgHead error ", err)
			break
		}
		if msgHead.GetMsgLen() > 0{
			msg := msgHead.(*Znet.Message)
			msg.Data = make([] byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil{
				fmt.Println("read  msg pdata error", err)
				break
			}
			fmt.Println("------> Recv Server Msg:ID =", msg.Id,
				", len: ", msg.DataLen, ", data:", string(msg.Data))
		}

		//cpu 阻塞
		time.Sleep(1*time.Second)
	}

}