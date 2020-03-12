package models

import (
	"encoding/json"
	"time"
)

//Rfqs --
type Rfqs struct {
	ID            int64 `gorm:"primary_key:true" json:"id"`
	TransactionID int64 `gorm:"column:transaction_id" json:"transaction_id"`
	RequestedBy   int64 `gorm:"column:requested_by" json:"requested_by"`
	CreatedBy     int64 `gorm:"column:created_by" json:"created_by"`

	PaymentType     int8  `gorm:"column:payment_type" json:"payment_type"`
	PaymentDuration int64 `gorm:"column:payment_duration" json:"payment_duration"`

	CompanyBuyerID  int64 `gorm:"column:company_buyer_id" json:"company_buyer_id"`
	CompanySellerID int64 `gorm:"column:company_seller_id" json:"company_seller_id"`

	RfqNo         string `gorm:"column:rfq_no" json:"rfq_no"`
	ReferenceNo   string `gorm:"column:reference_no" json:"reference_no"`
	Notes         string `gorm:"column:notes" json:"notes"`
	NoteForSeller string `gorm:"column:note_for_seller" json:"note_for_seller"`

	SubTotal   float64 `gorm:"column:sub_total" json:"sub_total"`
	TaxBasis   float64 `gorm:"column:tax_basis" json:"tax_basis"`
	Ppn        float64 `gorm:"column:ppn" json:"ppn"`
	Pph        float64 `gorm:"column:pph" json:"pph"`
	Rounding   float64 `gorm:"column:rounding" json:"rounding"`
	GrandTotal float64 `gorm:"column:grand_total" json:"grand_total"`

	Status       int8   `gorm:"column:status" json:"status"`
	StatusReason string `gorm:"column:status_reason" json:"status_reason"`

	ExpiredAt time.Time  `gorm:"column:expired_at" json:"expired_at"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

//MarshalJSON --
func (u *Rfqs) MarshalJSON() ([]byte, error) {
	type Alias Rfqs

	usr := &struct {
		Password string `json:"password,omitempty"`
		*Alias
	}{
		Password: "",
		Alias:    (*Alias)(u),
	}
	return json.Marshal(usr)
}
