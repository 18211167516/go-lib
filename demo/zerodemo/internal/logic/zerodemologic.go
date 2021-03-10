package logic

import (
	"context"

	"github.com/18211167516/go-lib/demo/zerodemo/internal/svc"
	"github.com/18211167516/go-lib/demo/zerodemo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ZerodemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZerodemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) ZerodemoLogic {
	return ZerodemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZerodemoLogic) Zerodemo(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
