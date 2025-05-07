package logic

import (
	"context"

	"zero-qp/panguService/rpc/internal/svc"
	"zero-qp/panguService/rpc/pangu"

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

func (l *PingLogic) Ping(in *pangu.Request) (*pangu.Response, error) {
	// todo: add your logic here and delete this line

	return &pangu.Response{}, nil
}
