package managers

import "gorm.io/gorm"

type ImageManager struct {
	gorm.Model
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&ImageManager{})
}

func (im *ImageManager) upload() {

}
