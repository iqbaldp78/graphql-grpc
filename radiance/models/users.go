package models

import (
	"encoding/json"
	"time"
)

//Users --
type Users struct {
	ID                   int64      `gorm:"primary_key:true" json:"id"`
	CompanyID            int64      `gorm:"column:company_id" json:"company_id"`
	Email                string     `gorm:"column:email" json:"email"`
	Password             string     `gorm:"column:password" json:"password"`
	VerifiedEmail        int8       `gorm:"column:verified_email" json:"verified_email"`
	Status               int8       `gorm:"column:status" json:"status"`
	Name                 string     `gorm:"column:name" json:"name"`
	Phone                string     `gorm:"column:phone" json:"phone"`
	Mobile               string     `gorm:"column:mobile" json:"mobile"`
	JobTitle             string     `gorm:"column:job_title" json:"job_title"`
	ImageStorageID       int64      `gorm:"column:image_storage_id" json:"image_storage_id"`
	LastLogin            *time.Time `gorm:"column:last_login_at" json:"last_login"`
	CreatedAt            time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt            *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"DEFAULT:NULL"`
	IsNotifPaymentReturn int8       `gorm:"column:is_notif_payment_return" json:"is_notif_payment_return"`
}

//MarshalJSON --
func (u *Users) MarshalJSON() ([]byte, error) {
	type Alias Users

	usr := &struct {
		Password string `json:"password,omitempty"`
		*Alias
	}{
		Password: "",
		Alias:    (*Alias)(u),
	}
	return json.Marshal(usr)
}
