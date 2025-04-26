package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
	"zero-qp/userService/rpc/userrpc"
	"zero-qp/wsService/api/internal/svc"
)

var (
	websocketUpgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type CheckOrigin func(r *http.Request) bool

type Manager struct {
	sync.RWMutex
	WebsocketUpgrader *websocket.Upgrader
	CheckOrigin       CheckOrigin
	Clients           map[string]Connection
	ReadMsgChan       chan *MsgPack
	Svc               *svc.ServiceContext
}

func NewManager(serviceContext *svc.ServiceContext) *Manager {
	return &Manager{Clients: make(map[string]Connection, 1024), ReadMsgChan: make(chan *MsgPack, 1024), Svc: serviceContext}
}

func (m *Manager) Run() {
	go m.ReadMsgPackChan()

	http.HandleFunc("/", m.serverWs)
	err := http.ListenAndServe("0.0.0.0:14000", nil)
	if err != nil {
		logx.Errorf("Run ListenAndServe err:%v", err)
	}

}

func (m *Manager) serverWs(w http.ResponseWriter, r *http.Request) {
	m.WebsocketUpgrader = &websocketUpgrader
	conn, err := m.WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Errorf("serverWs WebsocketUpgrader.Upgrade err: %v", err)
	}
	wsConnection := NewWsConnection(conn, m)
	//add
	m.AddClient(wsConnection)
	wsConnection.Run()

}

func (m *Manager) DelCilent(ws *WsConnection) {
	//删除cid
	for cid, c := range m.Clients {
		if cid == ws.Cid {
			delete(m.Clients, cid)
			c.Close()
		}
	}
}

func (m *Manager) AddClient(conn *WsConnection) {
	m.Lock()
	defer m.Unlock()
	m.Clients[conn.Cid] = conn

}

func (m *Manager) SendMsgAllUser() {
	for k, c := range m.Clients {
		fmt.Println(k)
		c.SendMsg([]byte(k))
	}

}
func (m *Manager) ReadMsgPackChan() {
	for {
		select {
		case msg, ok := <-m.ReadMsgChan:
			if ok {
				m.ReadMsgPackHander(msg)
			}
		}
	}
}

func (m *Manager) ReadMsgPackHander(msgPack *MsgPack) {

	bodyMap := make(map[string]string, 0)
	err := json.Unmarshal(msgPack.Body, &bodyMap)
	if err != nil {
		m.Clients[msgPack.Cid].SendMsg([]byte("请发送正确的数据！"))
	}

	msgType := bodyMap["msg_type"]

	switch msgType {
	case "1":
		bindUser := CidBindUser{}
		err = json.Unmarshal(msgPack.Body, &bindUser)
		if err != nil {
			m.Clients[msgPack.Cid].SendMsg([]byte("绑定错误"))
			return
		}
		resp, rpcErr := m.Svc.User.CidBindUid(context.Background(), &userrpc.CidBindUidReq{Uid: bindUser.Uid, Cid: msgPack.Cid})
		if rpcErr != nil || !resp.IsSuccess {
			m.Clients[msgPack.Cid].SendMsg([]byte("绑定错误"))
			return
		}

	default:
		m.Clients[msgPack.Cid].SendMsg([]byte("未知错误"))
	}

}
