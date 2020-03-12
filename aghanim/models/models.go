package models

import "github.com/jinzhu/gorm"

//Base --
type Base struct {
}

// Repository --
type Repository interface {
	ModelsGetRfqs(db *gorm.DB, req int) error
}
