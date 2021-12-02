package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//在完成了配置文件的确定和编写后，我们需要针对读取配置的行为进行封装，便于应用程序的使用

type Setting struct {
	vp *viper.Viper
}

// NewSetting 初始化本项目的配置的基础属性
// 设定配置文件的名称为 config，配置类型为 yaml，并且设置其配置路径为相对路径 configs/
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config) //可以设置多个配置路径,解决路径查找问题
		}
	}
	vp.SetConfigType("yaml") //设置配置文件类型
	err := vp.ReadInConfig() //加载配置文件
	if err != nil {
		return nil, err
	}
	s := &Setting{vp: vp}
	s.WatchSettingChange() //热监听
	return s, nil
}

//配置名存储记录
var sections = make(map[string]interface{})

// ReadSection 绑定配置文件
func (s *Setting) ReadSection(k string, v interface{}) error {
	//绑定
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// ReloadAllSection 重新读取配置文件
func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// WatchSettingChange 监听配置文件
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			err := s.ReloadAllSection()
			if err != nil {
				panic(err)
			}
		})
	}()
}
