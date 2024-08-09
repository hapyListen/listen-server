package songlist

import (
	"context"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取歌曲
func NewGetSongLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSongLogic {
	return &GetSongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSongLogic) GetSong(req *types.GetSongReq) (resp *types.GetSongResp, err error) {
	// todo: add your logic here and delete this line

	return
}
