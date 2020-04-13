package Znet

import (
	"awesomeProject/Zinx/Ziface"
	"awesomeProject/Zinx/utils"
	"fmt"
	"net"
)
// iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器的IP
	IP string
	//服务器的端口
	Port int
	//当前的Server添加一个router, server注册的连接对应的处理业务
	Router Ziface.IRouter
}



func (s *Server) Start(){
	fmt.Printf("[Struct] Server Name:%s Listenner at IP:%s, Port:%d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Struct] Server Listenner at IP:%s, Port:%d\n", s.IP, s.Port)

	go func(){
		// 1 获取一个TCP的addr
		addr, err :=net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d",s.IP, s.Port))
		if err != nil{
			fmt.Println("resolve tcp addr error:", err)
			return
		}

		// 2 尝试监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil{
			fmt.Println("listenner:",listener, "err:", err)
			return
		}
		fmt.Println("Start Zinx server succ, ", s.Name, ", Listenning..")
		var cid uint32;
		cid = 0

		// 3 阻塞的等待客户端连接，处理客户端的连接业务（读写）
		for{
			// 如果有客户端连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil{
				fmt.Println("Accept Error:",err)
				continue
			}

			// 已经与客户端简历连接，做一个最大512字节长度的回显业务

			dealConn := NewConnection(conn, cid, s.Router)
			cid ++

			//启动
			go dealConn.Start()
		}
	}()


}

func (s *Server) Stop(){
	//TODO 将一些服务器的资源，状态或者一些已经开辟的连接信息，进行停止或者回收
}

// 运行服务器
func (s *Server) Serve(){
	//启动Sever的服务器功能
	s.Start()
	//todo 做一些服务器启动之后的额外业务

	//阻塞状态
	select{}
}

// 添加Router
func (s *Server) AddRouter(router Ziface.IRouter){
	s.Router = router
}

/*
初始化服务器方法
 */
func NewServer(name string) Ziface.IServer{
	s := &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TcpPort,
		Router:nil,
	}
	return s
}
