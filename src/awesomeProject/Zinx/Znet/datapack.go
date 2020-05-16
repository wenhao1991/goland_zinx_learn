package Znet

import (
	"awesomeProject/Zinx/Ziface"
	"awesomeProject/Zinx/utils"
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
)

// 封包， 拆包的具体模块
type DataPack struct {

}

// 拆包封包实例的一个初始化方法
func NewDataPack() *DataPack{
	return &DataPack{}
}

func (d *DataPack)GetHeadLen() uint32{
	// DataLen uint32（4字节）+ ID uint32(4字节)
	return 8
}
//封包方法
// |datalen|msgId|data|
func (d *DataPack) Pack(msg Ziface.IMessage)([]byte, error){
	// 创建一个存放byte字节流的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	// 将dataLen写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil{
		return nil, err
	}
	//将MsgId写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil{
		return nil, err
	}
	//将data数据写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil{
		return nil, err
	}
	return dataBuff.Bytes(), nil
}
// 拆包方法 (将包的head信息读出来)， 之后再根据head信息里的data的长度，在进行一次读
func (d *DataPack) Unpack(binaryData []byte) (Ziface.IMessage, error){
	// 创建一个从输入二进制数据的ioReader
	dataBuff :=bytes.NewReader(binaryData)

	//只解压head信息，得到datalen和MsgId
	msg := &Message{}

	// 读dataLen
	if err :=binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err!=nil{
		return nil, err
	}
	// 读MsgId
	if err :=binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err!=nil{
		return nil, err
	}
	// 判断datalen是否已经超出了我们允许的最大包长度
	if (utils.GlobalObject.MaxPackageSize >0 && msg.DataLen > utils.GlobalObject.MaxPackageSize){
		return nil, errors.New("too Large msg data recv! msg_len:"+strconv.FormatInt(int64(msg.GetMsgLen()), 10) + " max:"+strconv.FormatInt(int64(utils.GlobalObject.MaxPackageSize), 10))
	}
	return msg, nil
}
