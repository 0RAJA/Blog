package middleware

import (
	"Blog/pkg/app"
	"Blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

//接口控制中间件

// JWT
/*
我们通过 GetHeader 方法从 Header 中获取 token 参数，
并调用 ParseToken 对其进行解析，再根据返回的错误类型进行断言判定
*/
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist { //查询token
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token) //解析token
			if err != nil {
				ecode = errcode.UnauthorizedTokenError
			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort() // 阻断
			return
		}
		c.Next() // 继续
	}
}
