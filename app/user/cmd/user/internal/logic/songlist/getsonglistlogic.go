package songlist

import (
	"context"

	"listen-server/app/user/cmd/user/internal/svc"
	"listen-server/app/user/cmd/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSongListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取歌单
func NewGetSongListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSongListLogic {
	return &GetSongListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSongListLogic) GetSongList(req *types.GetSongListReq) (resp *types.GetSongListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
