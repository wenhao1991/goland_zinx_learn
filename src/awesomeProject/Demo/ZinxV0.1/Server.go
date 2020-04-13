package main

import (
	"awesomeProject/Zinx/Ziface"
	"awesomeProject/Zinx/Znet"
	"fmt"
)
/*
	基于Zinx框架来开发的服务器应用程序
 */

// ping test 自定义路由
type PingRouter struct {
	Znet.BaseRouter
}


// Test PreHandle
func (this *PingRouter)PreHandle(request Ziface.IRequest){
	fmt.Println("Call Router PreHandle...")
	_, err :=request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil{
		fmt.Println("call back before ping error")
	}
}
// Test Handle
func (this *PingRouter)Handle(request Ziface.IRequest){
	fmt.Println("Call Router Handle...")
	_, err :=request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil{
		fmt.Println("call back ping...ping error")
	}
}

// Test PostHandle
func (this *PingRouter)PosHandle(request Ziface.IRequest){
	fmt.Println("Call Router PosHandle...")
	_, err :=request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil{
		fmt.Println("call back afterping error")
	}
}

func main(){
	//1 创建一个server 句柄
	s := Znet.NewServer("[zinx V0.1]")
	//2 给当前的znix 框架添加一个自定义的router
	s.AddRouter(&PingRouter{})
	s.Serve()
}
