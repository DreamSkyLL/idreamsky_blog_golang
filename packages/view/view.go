package view

import (
	"gorm.io/gorm"
)

type BasicView struct {
	Query *gorm.DB
}

func (v BasicView) GetObject() *gorm.DB {
	return v.Query
}

func (v BasicView) GetObjects() *gorm.DB {
	return v.Query
}

// func (vs BasicView) FilterObjects(*gorm.DB) *gorm.DB

// func (vs BasicView) GetPagination(*gorm.DB) *gorm.DB
