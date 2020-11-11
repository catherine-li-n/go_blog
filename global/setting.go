package global

import (
	"github.com/catherine.li/go_blog/pkg/logger"
	"github.com/catherine.li/go_blog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting    *setting.AppSettingS
	//EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
