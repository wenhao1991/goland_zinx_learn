package Ziface

/*
	IRequest接口
	实际上是吧客户端请求的连接和数据包装到一个Request中
 */

type IRequest interface {
	// 得到当前链接
	GetConnection() IConnection
	// 得到当前数据
	GetData() []byte
	GetMsgID() uint32
}
