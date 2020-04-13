package Znet

import "awesomeProject/Zinx/Ziface"

type Request struct {
	// 已经和客户端建立好的连接
	conn Ziface.IConnection
	// 客户端请求的数据
	data []byte
}

func(r Request) GetConnection() Ziface.IConnection{
	return r.conn
}
// 得到当前数据
func(r Request) GetData() []byte{
	return r.data
}