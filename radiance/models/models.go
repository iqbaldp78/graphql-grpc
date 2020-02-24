package models

import "github.com/jinzhu/gorm"

//Base --
type Base struct {
}

// Repository --
type Repository interface {
	ModelsGetAllUsers(db *gorm.DB, req int) error
}
