package posts

import "github.com/gin-gonic/gin"

func ArticleRegister(r *gin.RouterGroup) {
	r.GET("", ArticleList)
	r.GET("/:id", ArticleRetrieve)
	r.POST("", ArticleAdd)
	r.PATCH("/:id", ArticleUpdate)
	r.DELETE("/:id", ArticleDelete)
}

func RegisterRoute(r *gin.RouterGroup) {
	ArticleRegister(r.Group("/article"))
}
