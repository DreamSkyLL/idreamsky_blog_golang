package main

import (
	"idreamsky_blog/global"
	"idreamsky_blog/initializer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Init Redis

	// Init DB
	global.MY_DB = initializer.InitDB()

	// Init Gin
	r := gin.Default()
	initializer.RegisterRoute(r)
	r.Run(":9999")
}
