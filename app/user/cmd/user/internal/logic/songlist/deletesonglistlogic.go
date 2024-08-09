package songlist

import (
	"context"

	"listen-server/app/user/cmd/user/internal/svc"
	"listen-server/app/user/cmd/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSonglistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除歌单
func NewDeleteSonglistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSonglistLogic {
	return &DeleteSonglistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSonglistLogic) DeleteSonglist(req *types.DeleteSonglistReq) (resp *types.DeleteSonglistResp, err error) {
	// todo: add your logic here and delete this line

	return
}
