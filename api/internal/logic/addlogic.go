package logic

import (
	"context"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (resp *types.AddResp, err error) {
	// todo: add your logic here and delete this line

	return &types.AddResp{
		Ok:true,
	},nil
}
