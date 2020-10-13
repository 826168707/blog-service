package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)


var (
	// 关联配置属性结构体
	ServerSetting	*setting.ServerSettingS
	AppSetting		*setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	// 关联数据库DB
	DBEngine *gorm.DB

	// 关联日志
	Logger *logger.Logger
)


