package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" json:"app_key,omitempty" binding:"required"`
	AppSecret string `form:"app_secret" json:"app_secret,omitempty" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
