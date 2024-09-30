package svc

import (
	"ad.com/http/internal/config"
	"ad.com/pkg/repo/group"
	group_contact "ad.com/pkg/repo/group-contact"
	"ad.com/pkg/repo/user"
	user_contact "ad.com/pkg/repo/user-contact"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Conn              sqlx.SqlConn
	Config            config.Config
	UserModel         user.UserModel
	UserContactModel  user_contact.UserContactModel
	GroupModel        group.GroupModel
	GroupContactModel group_contact.GroupContactModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 这里可以控制每个不同的model可以有不同的数据库连接
	conn := sqlx.NewMysql(c.Datasource["Mysql"])
	return &ServiceContext{
		Conn:              conn,
		Config:            c,
		UserModel:         user.NewUserModel(conn),
		UserContactModel:  user_contact.NewUserContactModel(conn),
		GroupModel:        group.NewGroupModel(conn),
		GroupContactModel: group_contact.NewGroupContactModel(conn),
	}
}
