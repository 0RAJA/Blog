package api

import (
	"Blog/global"
	"Blog/internal/service"
	"Blog/pkg/app"
	"Blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// GetAuth
/*
这块的逻辑主要是校验及获取入参后，绑定并获取到的 app_key 和 app_secrect 进行数据库查询，
检查认证信息是否存在，若存在则进行 Token 的生成并返回
*/

// GetAuth
// @Tags Auth
// @Summary 获取Token
// @Produce  json
// @Param AppKey body string true "AppKey"
// @Param AppSecret body string true "AppSecret"
// @Success 200 {string} string "token"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param) //检查是否存在认证信息
	if err != nil {
		global.Logger.Infof("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret) //通过认证信息生成Token
	if err != nil {
		global.Logger.Infof("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
