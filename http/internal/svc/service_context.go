package svc

import (
	"ad.com/http/internal/config"
	"ad.com/pkg/repo/category"
	"ad.com/pkg/repo/group"
	group_contact "ad.com/pkg/repo/group-contact"
	"ad.com/pkg/repo/resource"
	"ad.com/pkg/repo/user"
	user_contact "ad.com/pkg/repo/user-contact"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	Conn              sqlx.SqlConn
	UserModel         user.UserModel
	UserContactModel  user_contact.UserContactModel
	GroupModel        group.GroupModel
	GroupContactModel group_contact.GroupContactModel
	CategoryModel     category.CategoryModel
	ResourceModel     resource.ResourceModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// here can make different model bind to different db
	conn := sqlx.NewMysql(c.DataSource["Mysql"])
	return &ServiceContext{
		Config:            c,
		Conn:              conn,
		UserModel:         user.NewUserModel(conn),
		UserContactModel:  user_contact.NewUserContactModel(conn),
		GroupModel:        group.NewGroupModel(conn),
		GroupContactModel: group_contact.NewGroupContactModel(conn),
		CategoryModel:     category.NewCategoryModel(conn),
		ResourceModel:     resource.NewResourceModel(conn),
	}
}
