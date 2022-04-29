package initializer

import (
	"idreamsky_blog/apps/posts"
	"idreamsky_blog/apps/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Connect to DB
	dsn := "root:ILovexixi0.@tcp(127.0.0.1:3306)/idreamsky_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate DB
	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {
	users.Migrate(db)
	posts.Migrate(db)
}
