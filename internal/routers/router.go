package routers

import (
	_ "Blog/docs" //记得导入文档的init函数
	"Blog/global"
	"Blog/internal/middleware"
	"Blog/internal/routers/api"
	v1 "Blog/internal/routers/api/v1"
	"Blog/pkg/limiter"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.Default()
	//接口文档绑定
	//初始化 docs 包和注册一个针对 swagger 的路由
	//而在初始化 docs 包后，其 swagger.json 将会默认指向当前应用所启动的域名下的 swagger/doc.json 路径
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Use(middleware.Recovery())                                              //防止崩溃
	r.Use(middleware.AppInfo())                                               //服务信息存储
	r.Use(middleware.Translations())                                          //翻译中间件
	r.Use(middleware.RateLimiter(methodLimiters))                             //限流桶
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout)) //超时控制
	r.Use(middleware.AccessLog())                                             //访问日志记录

	//上传文件绑定
	r.POST("/upload/file", NewUpload().UpLoadFile)        // 上传文件路由绑定
	r.Static("/static", global.AppSetting.UploadSavePath) // 静态资源文件绑定

	//获取token绑定
	r.POST("/auth", api.GetAuth)

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT()) //使用JWT
	{
		/*
			GET：读取/检索动作。
			POST：新增/新建动作。
			PUT：更新动作，用于更新一个完整的资源，要求为幂等。
			PATCH：更新动作，用于更新某一个资源的一个组成部分，也就是只需要更新该资源的某一项，就应该使用 PATCH 而不是 PUT，可以不幂等。
			DELETE：删除动作。
		*/
		apiv1.POST("tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
