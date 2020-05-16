package Znet

import (
	"awesomeProject/Zinx/Ziface"
	"errors"
	"fmt"
	"io"
	"net"
)
/*
 连接模块
 */
type Connection struct{
	// 当前链接的socket TCP 套接字
	Conn  *net.TCPConn

	// 链接的id
	ConnId uint32

	// 当前链接的状态
	isClosed bool

	// 告知当前链接已退出/停止的channel
	ExitChan chan bool

	//该连接处理的方法Router
	Router Ziface.IRouter
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router Ziface.IRouter) *Connection{
	c := &Connection{
		Conn: conn,
		ConnId:    connID,
		isClosed:  false,
		Router: router,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// 链接的读业务方法
func(c *Connection) StartReader(){
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connId= ", c.ConnId, " Reader is exit, remote addr is ", c.RemoteAddr())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中， 最大utils.GlobalObject.MaxPackageSize字节
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil{
		//	fmt.Println("recv buf err", err)
		//	continue
		//}
		// create a pack instance
		dp := NewDataPack()

		// read head
		headData := make([]byte, dp.GetHeadLen())
		if _, err :=io.ReadFull(c.GetTCPConnection(), headData); err != nil{
			fmt.Println("read head err:", err)
			break
		}
		// get msgId and  msg Datalen, put in msg
		msg, err := dp.Unpack(headData)
		if err != nil{
			fmt.Println("unpack err:", err)
			break
		}
		// get data upon Datalen, pu in msg
		var data []byte
		if msg.GetMsgLen() > 0{
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err!=nil{
				fmt.Println("read msg, data err", err)
				break
			}
		}
		msg.SetData(data)
		//得到当前的conn数据的Request的请求数据
		req := Request{conn:c, msg:msg}

		//从路由中，找到注册绑定的Conn对应的router调用
		go func(request Ziface.IRequest){
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PosHandle(request)
		}(&req)

		////调用当前链接所绑定的HandleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil{
		//	fmt.Println("ConnId ", c.ConnId, " handle is error", err)
		//	break
		//}
	}
}

func (c *Connection) Start(){
	fmt.Println("Conn Start().. ConnnId= ", c.ConnId)
	//启动从当前链接的读业务
	go c.StartReader()
	//TODO 启动从当前链接写数据的业务

}
//停止连接
func (c *Connection) Stop(){
	fmt.Println("Conn Stop().. ConnnId= ", c.ConnId)
	if c.isClosed == true{
		return
	}
	c.isClosed=true

	//关闭socket连接
	c.Conn.Close()
	close(c.ExitChan)
}
//获取当前绑定的conn
func (c *Connection) GetTCPConnection() *net.TCPConn{
	return c.Conn
}
//获取当前链接模块的链接id
func (c *Connection) GetConnID() uint32{
	return c.ConnId
}
//获取客户端tcp状态 ip port
func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}
//发送数据给客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error{
	if c.isClosed == true{
		return errors.New("Connection is close when send msg")
	}
	// pack data
	dp := NewDataPack()
	binaryMsg, err :=dp.Pack(NewMessage(msgId, data))
	if err != nil{
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg")
	}
	if _, err :=c.Conn.Write(binaryMsg); err!=nil{
		fmt.Println("Write msg id = ", msgId, " err:", err)
		return errors.New("Write msg err")
	}

	return nil
}