package initializer

import (
	"idreamsky_blog/apps/posts"
	"idreamsky_blog/apps/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	v := r.Group("/api")

	users.RegisterRoute(v.Group("/auth"))
	posts.RegisterRoute(v.Group("/post"))
}
