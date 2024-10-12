// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package svc

import (
	"ad.com/http/internal/config"
	"ad.com/pkg/repo/category"
	"ad.com/pkg/repo/group"
	"ad.com/pkg/repo/group_contact"
	"ad.com/pkg/repo/resource"
	"ad.com/pkg/repo/user"
	"ad.com/pkg/repo/user_contact"
	zeroRds "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	Conn              sqlx.SqlConn
	RdsClient         *zeroRds.Redis
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
	rdsClient := zeroRds.MustNewRedis(c.RedisConf)

	return &ServiceContext{
		Config:            c,
		Conn:              conn,
		RdsClient:         rdsClient,
		UserModel:         user.NewUserModel(conn),
		UserContactModel:  user_contact.NewUserContactModel(conn),
		GroupModel:        group.NewGroupModel(conn),
		GroupContactModel: group_contact.NewGroupContactModel(conn),
		CategoryModel:     category.NewCategoryModel(conn),
		ResourceModel:     resource.NewResourceModel(conn),
	}
}
