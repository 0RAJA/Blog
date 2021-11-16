package middleware

import (
	"Blog/global"
	"Blog/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

/*
访问日志记录
在出问题时，我们常常会需要去查日志，那么除了查错误日志、业务日志以外，还有一个很重要的日志类别，就是访问日志，
从功能上来讲，它最基本的会记录每一次请求的请求方法、方法调用开始时间、方法调用结束时间、方法响应结果、方法响应结果状态码，
更进一步的话，会记录 RequestId、TraceId、SpanId 等等附加属性，以此来达到日志链路追踪的效果，如下图：


但是在正式开始前，你又会遇到一个问题，你没办法非常直接的获取到方法所返回的响应主体，这时候我们需要巧妙利用 Go interface 的特性，
实际上在写入流时，调用的是 http.ResponseWriter，如下：

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}
那么我们只需要写一个针对访问日志的 Writer 结构体，实现我们特定的 Write 方法就可以解决无法直接取到方法响应主体的问题了
*/

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//我们在 AccessLogWriter 的 Write 方法中，实现了双写
//因此我们可以直接通过 AccessLogWriter 的 body 取到值
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWrite := &AccessLogWriter{body: new(bytes.Buffer), ResponseWriter: c.Writer}
		c.Writer = bodyWrite
		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(), //PostForm包含来自PATCH、POST或PUT主体参数的解析表单数据 当前的请求参数
			"response": bodyWrite.body.String(),     //当前的请求结果响应主体
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,   //当前的调用方法
			bodyWrite.Status(), //当前的响应结果状态码
			beginTime,          //调用方法的开始时间，调用方法结束的结束时间
			endTime,
		)
	}
}
