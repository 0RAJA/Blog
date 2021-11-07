package global

import "Blog/pkg/setting"

//我们需要将配置信息和应用程序关联起来

var (
	ServerSetting   = new(setting.ServerSettingS)
	AppSetting      = new(setting.AppSettingS)
	DatabaseSetting = new(setting.DatabaseSettingS)
)

