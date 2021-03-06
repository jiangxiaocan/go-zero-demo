package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-demo/demomall/service/user/api/internal/logic"
	"go-zero-demo/demomall/service/user/api/internal/svc"
	"go-zero-demo/demomall/service/user/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
