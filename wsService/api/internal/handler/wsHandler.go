package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-qp/wsService/api/internal/logic"
	"zero-qp/wsService/api/internal/net"
	"zero-qp/wsService/api/internal/svc"
)

func wsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		net.NetManager.SendMsgAllUser()
		l := logic.NewWsLogic(r.Context(), svcCtx)
		err := l.Ws()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
