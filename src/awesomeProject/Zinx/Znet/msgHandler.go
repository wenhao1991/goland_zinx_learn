package Znet

import (
	"awesomeProject/Zinx/Ziface"
	"fmt"
	"strconv"
)

/*
	消息处理模块的实现
 */

type MsgHandler struct {

	Apis map[uint32] Ziface.IRouter
}

// 初始化/创建MsgHandler方法

func NewMsgHandler() *MsgHandler{
	return &MsgHandler{
		Apis:make(map[uint32] Ziface.IRouter),
	}
}

func (mh* MsgHandler)DoMsgHandler(request Ziface.IRequest){
	//1 从Request中找到msgID
	handler, ok :=mh.Apis[request.GetMsgID()]
	if !ok{
		fmt.Println("api msgID = ", request.GetMsgID(), "is NOT FOUND! Need Register!")
		return
	}
	// 2 根据MsgID 调度对应router业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PosHandle(request)
}

//为消息添加具体的处理逻辑
func (mh* MsgHandler)AddRouter(msgID uint32, router Ziface.IRouter){
	//1 判断当前msg办公的处理方法是否已经存在
	if _, ok:= mh.Apis[msgID]; ok{
		//id已注册了
		panic("repeat api, msgID=" + strconv.Itoa(int(msgID)))
	}
	//2 添加msg与API绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID =", msgID, " succ!")
}