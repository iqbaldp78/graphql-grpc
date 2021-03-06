// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Rfqs struct {
	ID              int        `json:"ID"`
	TransactionID   int        `json:"TransactionID"`
	RequestedBy     int        `json:"RequestedBy"`
	CreatedBy       int        `json:"CreatedBy"`
	PaymentType     int        `json:"PaymentType"`
	PaymentDuration int        `json:"PaymentDuration"`
	CompanyBuyerID  int        `json:"CompanyBuyerID"`
	CompanySellerID int        `json:"CompanySellerID"`
	RfqNo           string     `json:"RfqNo"`
	ReferenceNo     string     `json:"ReferenceNo"`
	Notes           string     `json:"Notes"`
	NoteForSeller   string     `json:"NoteForSeller"`
	SubTotal        float64    `json:"SubTotal"`
	TaxBasis        float64    `json:"TaxBasis"`
	Ppn             float64    `json:"Ppn"`
	Pph             float64    `json:"Pph"`
	Rounding        float64    `json:"Rounding"`
	GrandTotal      float64    `json:"GrandTotal"`
	Status          int        `json:"Status"`
	StatusReason    string     `json:"StatusReason"`
	ExpiredAt       *time.Time `json:"ExpiredAt"`
	CreatedAt       *time.Time `json:"CreatedAt"`
	UpdatedAt       *time.Time `json:"UpdatedAt"`
	DeletedAt       *time.Time `json:"DeletedAt"`
}
