package middleware

//链路追踪中间件

//func Tracing() func(c gin.Context) {
//	return func(c gin.Context) {
//		var newCtx context.Context
//		var span opentracing.Span
//		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
//		if err!=nil{
//
//		}
//	}
//}
