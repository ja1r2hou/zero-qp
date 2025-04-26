package logic

import (
	"context"

	"zero-qp/userService/rpc/internal/svc"
	"zero-qp/userService/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CidBindUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userSvc *svc.UserSvc
}

func NewCidBindUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CidBindUidLogic {
	return &CidBindUidLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userSvc: svc.NewUserSvc(ctx, svcCtx),
	}
}

// CidBindUid websocket cid绑定uid
func (l *CidBindUidLogic) CidBindUid(in *user.CidBindUidReq) (*user.CidBindUidResp, error) {
	return l.userSvc.CidBindUid(in)
}
