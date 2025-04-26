package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"zero-qp/common/jwts"
	"zero-qp/userService/rpc/userrpc"

	"zero-qp/gateService/api/internal/svc"
	"zero-qp/gateService/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {

	registerParams := &userrpc.RegisterParams{Account: req.Account, Password: req.Password, SmsCode: "", LoginPlatform: req.LoginPlatform}
	response, err := l.svcCtx.User.Register(context.Background(), registerParams)
	if err != nil {
		return nil, err
	}

	claims := jwts.CustomClaims{
		Uid: response.Uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	token, err := jwts.GenToken(&claims, l.svcCtx.Config.Jwt.Secret)
	if err != nil {
		logx.Error("Register jwt gen token err:%v", err)
	}
	registerResp := &types.RegisterResp{ServerInfo: types.ServerInfo{Host: "127.0.0.1", Port: 12000}, Token: token}

	return registerResp, nil
}
