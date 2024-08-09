package svc

import (
	"listen-server/app/user/cmd/api/user/internal/config"
	"listen-server/app/user/model/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
