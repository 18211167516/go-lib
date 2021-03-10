package handler

import (
	"net/http"

	"github.com/18211167516/go-lib/demo/zerodemo/internal/logic"
	"github.com/18211167516/go-lib/demo/zerodemo/internal/svc"
	"github.com/18211167516/go-lib/demo/zerodemo/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ZerodemoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewZerodemoLogic(r.Context(), ctx)
		resp, err := l.Zerodemo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
