package routers

import (
	v1 "Blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
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
