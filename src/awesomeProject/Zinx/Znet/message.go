package Znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(id uint32, data []byte) *Message{
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//获取消息Id
func (m *Message)GetMsgId()	uint32{
	return m.Id
}

func (m *Message)GetMsgLen() uint32{
	return m.DataLen
}

func (m *Message)GetData()	[]byte{
	return m.Data
}

//获取消息Id
func (m *Message)SetMsgId(id uint32){
	m.Id = id
}

func (m *Message)SetMsgLen(len uint32){
	m.DataLen = len
}

func (m *Message)SetData(data []byte){
	m.Data = data
}
