package net

type Connection interface {
	Close()
	SendMsg(msg []byte)
}

type MsgPack struct {
	Cid  string `json:"cid"`
	Body []byte `json:"body"`
}

type CidBindUser struct {
	Uid     string `json:"uid"`
	MsgType string `json:"msg_type"`
}
