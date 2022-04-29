package posts

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
}

type Column struct {
	gorm.Model
}

type Album struct {
	gorm.Model
}

type Moment struct {
	gorm.Model
}

type Article struct {
	gorm.Model
	Title      string
	Content    string
	Summary    string
	Visibility int8
	Deleted    bool
}

type Picture struct {
	gorm.Model
}
