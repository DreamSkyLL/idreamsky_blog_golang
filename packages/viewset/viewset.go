package viewset

import (
	"gorm.io/gorm"
)

type ViewSet struct {
	query *gorm.DB
}

func (vs ViewSet) GetObject() *gorm.DB {
	return vs.query
}

func (vs ViewSet) GetObjects() *gorm.DB

func (vs ViewSet) FilterObjects(*gorm.DB) *gorm.DB

func (vs ViewSet) GetPagination(*gorm.DB) *gorm.DB
