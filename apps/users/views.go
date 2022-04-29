package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.String(http.StatusOK, "Register")
}

func Login(c *gin.Context) {
	c.String(http.StatusOK, "Login")
}

func Retrieve(c *gin.Context) {

}

func UserGroupAdd(c *gin.Context) {

}

func UserGroupList(c *gin.Context) {

}

func UserGroupRetrieve(c *gin.Context) {

}

func UserGroupUpdate(c *gin.Context) {

}
