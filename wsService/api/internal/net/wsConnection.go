package net

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

var (
	pongWait = 10 * time.Second
	pingWait = 10 * time.Second
)

type WsConnection struct {
	Cid          string
	Manager      *Manager
	WsConn       *websocket.Conn
	ReadMsgChan  chan *MsgPack
	WriteMsgChan chan []byte
	pingTicker   *time.Ticker
}

func (c *WsConnection) SendMsg(msg []byte) {
	c.WsConn.WriteMessage(websocket.TextMessage, msg)

}

func (c *WsConnection) Close() {
	if c.WsConn != nil {
		logx.Errorf("close err:%v", c.WsConn.Close())
	}

}

func NewWsConnection(wsConn *websocket.Conn, manager *Manager) *WsConnection {
	return &WsConnection{WsConn: wsConn, Manager: manager, WriteMsgChan: make(chan []byte, 1024), Cid: uuid.New().String(), ReadMsgChan: manager.ReadMsgChan}
}

func (c *WsConnection) Run() {
	go c.ReadMsg()
	go c.WriteMsg()
	c.WsConn.SetPongHandler(c.PongHandler)

}

func (c *WsConnection) ReadMsg() {

	defer func() {
		c.Manager.DelCilent(c)
		c.Close()

	}()
	c.PongHandler("")
	c.WsConn.SetReadLimit(1024)
	for {
		messageType, message, err := c.WsConn.ReadMessage()
		if err != nil {
			logx.Errorf("ReadMsg ReadMessage err:%v", err)
			break
		}
		if messageType == websocket.TextMessage {
			if c.ReadMsgChan != nil {
				c.ReadMsgChan <- &MsgPack{Cid: c.Cid, Body: message}

			} else {
				logx.Errorf("ReadMsg not websocket.TextMessage msgType:%v ,msg:%v", messageType, string(message))
			}

		}
	}

}

func (c *WsConnection) WriteMsg() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case wmsg, ok := <-c.WriteMsgChan:
			if !ok {
				err := c.WsConn.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					logx.Errorf("WriteMsg WriteMsgChan not ok err:%v", err)
				}
			}

			err := c.WsConn.WriteMessage(websocket.TextMessage, wmsg)
			if err != nil {
				logx.Errorf("WriteMsg WriteMsgChan not ok err:%v", err)
			}
		case <-ticker.C:
			if len(c.Manager.Clients) > 0 {
				err := c.WsConn.SetWriteDeadline(time.Now().Add(pingWait))
				if err != nil {
					logx.Errorf("WriteMsg SetWriteDeadline err:%v", err)
				}

				err = c.WsConn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					logx.Errorf("WriteMsg PingMessage err:%v", err)
					c.WsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					return
				}
			}

		}
	}

}

func (c *WsConnection) PongHandler(data string) error {
	err := c.WsConn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		logx.Errorf("PongHandler SetReadDeadline err:%v", err)
		return err
	}
	return nil
}
