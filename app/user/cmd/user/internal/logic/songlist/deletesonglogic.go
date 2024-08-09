package songlist

import (
	"context"

	"listen-server/app/user/cmd/user/internal/svc"
	"listen-server/app/user/cmd/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除歌曲
func NewDeleteSongLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSongLogic {
	return &DeleteSongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSongLogic) DeleteSong(req *types.DeleteSongReq) (resp *types.DeleteSongResp, err error) {
	// todo: add your logic here and delete this line

	return
}
