package logic

import (
	"context"

	"zero-qp/payService/rpc/internal/svc"
	"zero-qp/payService/rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *pay.Request) (*pay.Response, error) {
	// todo: add your logic here and delete this line

	return &pay.Response{}, nil
}
