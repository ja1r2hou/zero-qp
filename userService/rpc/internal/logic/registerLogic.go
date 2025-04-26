package logic

import (
	"context"

	"zero-qp/userService/rpc/internal/svc"
	"zero-qp/userService/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userSvc *svc.UserSvc
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userSvc: svc.NewUserSvc(ctx, svcCtx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterParams) (*user.RegisterResponse, error) {

	return l.userSvc.WxRegister(l.ctx, in)
}
