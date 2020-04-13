package Ziface

import "net"

type IConnection interface {
	//启动连接 让当前的连接准备开始工作
	Start()
	//停止连接
	Stop()
	//获取当前绑定的conn
	GetTCPConnection() *net.TCPConn
	//获取当前链接模块的链接id
	GetConnID() uint32
	//获取客户端tcp状态 ip port
	RemoteAddr() net.Addr
	//发送数据给客户端
	Send(data []byte) error
}

//定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error