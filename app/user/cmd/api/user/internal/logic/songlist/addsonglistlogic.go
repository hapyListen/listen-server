package songlist

import (
	"context"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSonglistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSonglistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSonglistLogic {
	return &AddSonglistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSonglistLogic) AddSonglist(req *types.AddSonglistReq) (resp *types.AddSonglistResp, err error) {
	// todo: add your logic here and delete this line

	return
}
