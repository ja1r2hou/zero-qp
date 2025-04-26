package respx

import (
	"context"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type RespOk struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"traceId"`
}

type RespErr struct {
	Status  int    `json:"status"`
	Msg     string `json:"msg"` //错误码对应key
	TraceId string `json:"traceId"`
}

// OkJsonCtx writes v into w with 200 OK.
func OkJsonCtx(ctx context.Context, w http.ResponseWriter, v interface{}) {
	r := &RespOk{
		Status:  200,
		Data:    v,
		TraceId: trace.TraceIDFromContext(ctx),
	}
	output, _ := sonic.MarshalString(r)
	logx.Info("result:                " + output + "              ")
	WriteJson(w, http.StatusOK, r)
}

// ErrorCtx writes err into w.
func ErrorCtx(ctx context.Context, w http.ResponseWriter, err error) {
	r := &RespErr{
		Status:  500,
		Msg:     err.Error(),
		TraceId: trace.TraceIDFromContext(ctx),
	}
	output, _ := sonic.MarshalString(r)
	logx.Info("result:                " + output + "              ")
	WriteJson(w, http.StatusInternalServerError, r)
}

// WriteJson writes v as json string into w with code.
func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set(httpx.ContentType, httpx.JsonContentType)
	w.WriteHeader(code)

	if bs, err := json.Marshal(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
}
