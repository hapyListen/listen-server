package usermanager

import (
	"context"
	"fmt"
	"net/http"

	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"
	"listen-server/app/user/model/user"
	"listen-server/common/aes"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	if req.UserId == 0 || req.Name == "" || req.Password == "" {
		return nil, errors.New(http.StatusBadRequest, "userId == 0 or name == '' or password == '' ")
	}

	passwd, err := aes.EncryptAES([]byte(req.Password))
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "password format error")
	}
	insertInfo := user.User{
		UserId:          int64(req.UserId),
		Password:        passwd,
		Name:            req.Name,
		PhoneNumber:     req.PhoneNumber,
		Email:           req.Eamil,
		Avatar:          req.Avatar,
		PersonSignature: req.PersonSignature,
	}
	_, err = l.svcCtx.UserModel.Insert(context.Background(), &insertInfo)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("database operation failed, error %v", err))
	}

	return &types.UserRegisterResp{
		Base: types.Base{
			Code: http.StatusOK,
			Msg:  "OK",
		},
	}, nil
}
