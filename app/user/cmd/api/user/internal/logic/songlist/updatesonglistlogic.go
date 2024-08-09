package songlist

import (
	"context"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSongListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新歌单
func NewUpdateSongListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSongListLogic {
	return &UpdateSongListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSongListLogic) UpdateSongList(req *types.UpdateSongListReq) (resp *types.UpdateSongListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
