package global

import (
	"github.com/ruixijiejie/blog-service/pkg/logger"
	"github.com/ruixijiejie/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
