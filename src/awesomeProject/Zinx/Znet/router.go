package Znet
import "awesomeProject/Zinx/Ziface"

// 实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写就好了
type BaseRouter struct {}

// 有的Router不希望有PreHandle、PostHandle这俩个业务
// 在处理conn业务之前的钩子方法Hook
func (br *BaseRouter)PreHandle(request Ziface.IRequest){}

// 在处理conn业务的主钩子方法Hook
func (br *BaseRouter)Handle(request Ziface.IRequest){}

// 在处理conn业务之后的钩子方法Hook
func (br *BaseRouter)PosHandle(request Ziface.IRequest){}