package usermanager

import (
	"context"
	"fmt"
	"net/http"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"
	"listen-server/app/user/common/encrypt"
	"listen-server/common/jwt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	if req.UserId == 0 || req.Password == "" {
		return nil, errors.New(http.StatusBadRequest, "userId == 0 or password == '' ")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUserId(context.Background(), int64(req.UserId))
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("database operation failed, userId %v is not register %v", req.UserId))
	}
	passwd, err := encrypt.DecryptAES(userInfo.Password)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("database password format error, %v", err))
	}
	if string(passwd) != req.Password {
		return nil, errors.New(http.StatusBadRequest, "incorrect password")
	}

	// Generate JWT
	authCfg := l.svcCtx.Config.Auth
	token, err := jwt.GenerateJWT(authCfg.AccessSecret, authCfg.AccessExpire, req.UserId)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("GenerateJWT failed, %v", err))
	}

	return &types.UserLoginResp{
		Base: types.Base{
			Code: http.StatusOK,
			Msg:  "OK",
		},
		Token: token,
	}, nil
}
