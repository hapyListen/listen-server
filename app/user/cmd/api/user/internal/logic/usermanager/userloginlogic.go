package usermanager

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"
	"listen-server/common/aes"
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

// Generate JWT
func (l *UserLoginLogic) generateJWT(oldToken string, userId int) (token string, err error) {
	if oldToken != "" {
		userJwt, err := jwt.ParseJWT(l.svcCtx.Config.Auth.AccessSecret, oldToken)
		if err != nil {
			return "", fmt.Errorf("ParseJWT failed, %v", err)
		}

		exp, err := userJwt.GetExpirationTime()
		if err != nil {
			return "", fmt.Errorf("GetExpirationTime failed, %v", err)
		}
		if time.Now().Sub(exp.Time) < 0 {
			return oldToken, nil
		}
	}

	authCfg := l.svcCtx.Config.Auth
	token, err = jwt.GenerateJWT(authCfg.AccessSecret, authCfg.AccessExpire, userId)
	if err != nil {
		return "", errors.New(http.StatusBadRequest, fmt.Sprintf("GenerateJWT failed, %v", err))
	}

	return token, nil
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	if req.UserId == 0 || req.Password == "" {
		return nil, errors.New(http.StatusBadRequest, "userId == 0 or password == '' ")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUserId(context.Background(), int64(req.UserId))
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("database operation failed, userId %v is not register %v", req.UserId))
	}
	passwd, err := aes.DecryptAES(userInfo.Password)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("database password format error, %v", err))
	}
	if string(passwd) != req.Password {
		return nil, errors.New(http.StatusBadRequest, "incorrect password")
	}

	token, err := l.generateJWT(userInfo.Token, int(userInfo.UserId))
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("generateJWT failed, %v", err))
	}
	userInfo.Token = token

	if err = l.svcCtx.UserModel.Update(context.Background(), userInfo); err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("Update user info failed, %v", err))
	}

	return &types.UserLoginResp{
		Base: types.Base{
			Code: http.StatusOK,
			Msg:  "OK",
		},
		Token: token,
	}, nil
}
