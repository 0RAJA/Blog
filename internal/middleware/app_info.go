package middleware

import "github.com/gin-gonic/gin"

/*
平时我们经常会需要在进程内上下文设置一些内部信息，
例如是应用名称和应用版本号这类基本信息，也可以是业务属性的信息存储，
例如是根据不同的租户号获取不同的数据库实例对象，
这时候就需要有一个统一的地方处理
*/

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
