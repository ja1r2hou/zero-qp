package handler

import (
	"net/http"
	"zero-qp/common/biz"
	"zero-qp/common/respx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-qp/gateService/api/internal/logic"
	"zero-qp/gateService/api/internal/svc"
	"zero-qp/gateService/api/internal/types"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			respx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			respx.ErrorCtx(r.Context(), w, biz.Fail)
		} else {
			respx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
