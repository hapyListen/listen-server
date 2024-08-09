package songlist

import (
	"context"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSonglistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建歌单
func NewCreateSonglistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSonglistLogic {
	return &CreateSonglistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSonglistLogic) CreateSonglist(req *types.CreateSonglistReq) (resp *types.CreateSonglistResp, err error) {
	// todo: add your logic here and delete this line

	return
}
