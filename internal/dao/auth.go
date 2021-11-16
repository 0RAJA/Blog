package dao

import "Blog/internal/model"

func (d *Dao) GetAuth(appkey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appkey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}
