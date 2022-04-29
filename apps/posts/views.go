package posts

import (
	"idreamsky_blog/global"
	"idreamsky_blog/packages/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleList(c *gin.Context) {
	view := view.BasicView{
		Query: global.MY_DB.Find(&Article{}),
	}
	result := []Article{}
	view.GetObjects().Find(&result)
	c.JSON(http.StatusOK, result)
}

func ArticleAdd(c *gin.Context) {
	c.String(http.StatusOK, "Add Article")
}

func ArticleRetrieve(c *gin.Context) {
	c.String(http.StatusOK, "Retrieve Article")
}

func ArticleUpdate(c *gin.Context) {
	c.String(http.StatusOK, "Update Article")
}

func ArticleDelete(c *gin.Context) {
	c.String(http.StatusOK, "Delete Article")
}
