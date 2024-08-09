package songlist

import (
	"context"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加歌曲
func NewAddSongLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSongLogic {
	return &AddSongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSongLogic) AddSong(req *types.AddSongReq) (resp *types.AddSongResp, err error) {
	// todo: add your logic here and delete this line

	return
}
