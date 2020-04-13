package utils

import (
	"awesomeProject/Zinx/Ziface"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	存储一切有关的全局参数，供其他模块使用
	一些参数可以通过zinx.json 供用户配置
 */


type GlobalObj struct {
	/*
		Server
	 */
	TcpServer Ziface.IServer	//当前全局Server对象
	Host string					// 当前服务器监听的IP
	TcpPort int 				// 监听端口号
	Name string					// 当前服务器名称

	/*
		Zinx
	 */
	Version string 				// 当前版本号
	MaxConn int					// 当前服务器主机允许的最大连接数
	MaxPackageSize uint32		//当前框架数据包最大值
}

/*
 全局对外的Globalobj
 */

var GlobalObject *GlobalObj

/*
	从zinx.json取加载自定义参数
 */
func (g *GlobalObj) Reload(){
	dir,_ := os.Getwd()
	fmt.Printf("cur dir: %s\n", dir)
	data, err := ioutil.ReadFile("src/awesomeProject/Demo/ZinxV0.1/conf/zinx.json")
	if err != nil{
		panic(err)
	}
	//将json文件数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil{
		panic(err)
	}
}


/*
 give init func, init current GlobalObject
 */
func init(){
	// 如果配置文件没有加载， 默认值
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		TcpPort:        8843,
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.Reload()
}