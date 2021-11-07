package setting

import "github.com/spf13/viper"

//在完成了配置文件的确定和编写后，我们需要针对读取配置的行为进行封装，便于应用程序的使用

type Setting struct {
	vp *viper.Viper
}

// NewSetting 初始化本项目的配置的基础属性
// 设定配置文件的名称为 config，配置类型为 yaml，并且设置其配置路径为相对路径 configs/
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/") //可以设置多个配置路径,解决路径查找问题
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}
