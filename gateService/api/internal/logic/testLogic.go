package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-qp/gateService/api/internal/svc"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() error {
	// todo: add your logic here and delete this line

	return nil
}
