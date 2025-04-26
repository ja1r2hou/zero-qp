package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"zero-qp/common/biz"
	"zero-qp/common/cacheKey"
	"zero-qp/userService/rpc/internal/model"
	"zero-qp/userService/rpc/user"
	"zero-qp/userService/rpc/userrpc"
)

const Prefix = "QP"
const AccountIdRedisKey = "AccountId"
const AccountIdBegin = 10000

type UserSvc struct {
	logx.Logger
	ctx    context.Context
	svcCtx *ServiceContext
}

func NewUserSvc(ctx context.Context, svcCtx *ServiceContext) *UserSvc {
	return &UserSvc{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (u *UserSvc) WxRegister(ctx context.Context, req *user.RegisterParams) (*userrpc.RegisterResponse, error) {
	response := &userrpc.RegisterResponse{}
	account := &model.Account{
		WxAccount:  req.Account,
		CreateTime: time.Now(),
	}
	uid, err := u.NextAccountId()
	if err != nil {
		return response, biz.SqlError
	}
	account.Uid = uid
	response.Uid = uid
	table := u.svcCtx.Mongo.Collection("account")
	_, err = table.InsertOne(ctx, account)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (u *UserSvc) NextAccountId() (string, error) {
	//自增 给一个前缀
	return u.incr(Prefix + ":" + AccountIdRedisKey)
}

func (u *UserSvc) incr(key string) (string, error) {
	//判断此key是否存在 不存在 set 存在就自增
	todo := context.TODO()
	var exist int64
	var err error
	//0 代表不存在
	if u.svcCtx.Redis != nil {
		exist, err = u.svcCtx.Redis.Exists(todo, key).Result()
	}
	if exist == 0 {
		//不存在
		if u.svcCtx.Redis != nil {
			err = u.svcCtx.Redis.Set(todo, key, AccountIdBegin, 0).Err()
		}
		if err != nil {
			return "", err
		}
	}
	var id int64
	if u.svcCtx.Redis != nil {
		id, err = u.svcCtx.Redis.Incr(todo, key).Result()
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func (u *UserSvc) CidBindUid(req *user.CidBindUidReq) (*user.CidBindUidResp, error) {

	key := cacheKey.WsUidBind + req.Uid
	bytes, _ := json.Marshal(req)

	err := u.svcCtx.Redis.Set(u.ctx, key, string(bytes), 0).Err()
	if err != nil {
		u.Logger.Errorf("CidBindUid Redis.Set err: %v", err)
		return &user.CidBindUidResp{IsSuccess: false}, err
	}
	return &user.CidBindUidResp{IsSuccess: true}, nil

}

func (u *UserSvc) CidUnBindUid(req *user.CidBindUidReq) (*user.CidBindUidResp, error) {
	key := cacheKey.WsUidBind + req.Uid

	err := u.svcCtx.Redis.Del(u.ctx, key).Err()
	if err != nil {
		u.Logger.Errorf("CidUnBindUid Redis.del err: %v", err)
		return &user.CidBindUidResp{IsSuccess: false}, err
	}
	return &user.CidBindUidResp{IsSuccess: true}, nil

}
