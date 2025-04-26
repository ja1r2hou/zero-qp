package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-qp/wsService/api/internal/svc"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws() error {
	// todo: add your logic here and delete this line

	return nil
}
