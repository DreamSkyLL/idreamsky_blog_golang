package users

import "github.com/gin-gonic/gin"

func UserAuthenticationRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/retrieve", Retrieve)
}

func UserGroupRegister(router *gin.RouterGroup) {
	router.GET("", UserGroupList)
	router.GET("/:id", UserGroupRetrieve)
	router.POST("", UserGroupAdd)
	router.PATCH("/:id", UserGroupUpdate)
}

func UserPermissionRegister(router *gin.RouterGroup) {
	router.GET("", UserGroupList)
	router.GET("/:id", UserGroupRetrieve)
	router.POST("", UserGroupAdd)
	router.PATCH("/:id", UserGroupUpdate)
}

func RegisterRoute(r *gin.RouterGroup) {
	UserAuthenticationRegister(r.Group(""))
	UserGroupRegister(r.Group("/group"))
	UserPermissionRegister(r.Group("/permissions"))
}
