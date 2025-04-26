package net

import "zero-qp/wsService/api/internal/svc"

var NetManager *Manager

func InitWs(serviceContext *svc.ServiceContext) {

	manager := NewManager(serviceContext)
	NetManager = manager
	manager.Run()
}
