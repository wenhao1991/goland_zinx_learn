package Znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// 只是负责测试datapack拆包 封包的单元测试
func TestDataPack(t *testing.T){
	/*
		模拟服务器
	 */
	// 1、 创建socketTCP
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil{
		fmt.Println("server listen err: ", err)
		return
	}

	// 创建一个go负责从客户端处理业务
	go func(){
		// 2 从客户端读取数据拆包处理
		for {
			conn,err := listener.Accept()
			if err != nil{
				fmt.Println("server accept error", err)
			}

			go func(conn net.Conn){
				//处理客户端的请求
				//------拆包过程 ------
				//定义一个拆包对象dp
				dp := NewDataPack()
				for{
					// 1 read from conn, read head
					headData := make([]byte, dp.GetHeadLen())
					_, err :=io.ReadFull(conn, headData)	// read full
					if err != nil{
						fmt.Println("read head error", err)
						return
					}
					msgHead, err :=dp.Unpack(headData)
					if err != nil{
						fmt.Println("server unpack error", err)
						return
					}

					if msgHead.GetMsgLen() > 0 {
						//2 read from conn, read data upon datalen
						msg := msgHead.(* Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil{
							fmt.Println("server unpack data ,error", err)
							return
						}
						// has read full msg
						fmt.Println("----- Recv MsgID: ", msg.Id, " DataLen:", msg.DataLen, " Data:", string(msg.Data))

					}
				}

			}(conn)
		}
	}()

	/*
	   similate a client
	 */
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil{
		fmt.Println("client dial error ,error", err)
		return
	}
	// create a packdata entity dp
	dp := NewDataPack()

	// simulate two pack send
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'1', '1','1','1'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil{
		fmt.Println("client pack msg1 error ,error", err)
		return
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'2', '2','2', '2', '2','2','1'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil{
		fmt.Println("client pack msg2 error ,error", err)
		return
	}
	sendData := append(sendData1, sendData2...)
	conn.Write(sendData)
	// 客户端阻塞
	select {}

}