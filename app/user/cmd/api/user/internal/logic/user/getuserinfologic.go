package user

import (
	"context"
	"fmt"
	"net/http"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

// 1:听歌中~; 2:活跃者; 3:默默无闻的点歌达人
const (
	ListenStatusId  = 1
	ActiveStatusId  = 2
	ReleaseStatusId = 3
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TODO 后续升级映射关系
func getStatusFlag(flag int64) string {
	switch flag {
	case ListenStatusId:
		return "听歌中"
	case ActiveStatusId:
		return "活跃中"
	case ReleaseStatusId:
		return "默默无闻的点歌达人"
	}
	return "一无所有"
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUSerInfoReq) (resp *types.GetUSerInfoResp, err error) {
	if req.UserId == 0 {
		return nil, errors.New(http.StatusBadRequest, "userId == 0")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUserId(context.Background(), int64(req.UserId))
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("User query failed, %v", err))
	}

	return &types.GetUSerInfoResp{
		Name:            userInfo.Name,
		PhoneNunber:     userInfo.PhoneNumber,
		Eamil:           userInfo.Email,
		StatusFlag:      getStatusFlag(userInfo.StatusFlag),
		Avatar:          userInfo.Avatar,
		PersonSignature: userInfo.PersonSignature,
	}, nil
}
