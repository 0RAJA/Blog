package middleware

import (
	"Blog/global"
	"Blog/pkg/app"
	"Blog/pkg/email"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//异常捕获处理

//我们需要针对我们的公司内部情况或生态圈定制 Recovery 中间件，确保异常在被正常捕抓之余，要及时的被识别和处理
//自定义 Recovery

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() { //捕获错误
			if err := recover(); err != nil {
				global.Logger.WithCallsFrames().Infof("panic recover err: %v", err)

				err := defailtMailer.SendMail( //短信通知
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %s,%d", app.GetNowDateTimeStr(), time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Infof("mail.SendMail err: %v", err)
				}
				c.Abort() //阻断
			}
		}()
		c.Next()
	}
}
