package entities

import "github.com/gofrs/uuid"

type AddressCustomer struct {
	BaseModel
	CustomerID    uuid.UUID `json:"customer_id"`
	Customer      *Customer `gorm:"foreignKey:CustomerID"`
	StreetAddress string    `json:"street_address" gorm:"column:street_address"`
	SubDistrict   string    `json:"sub_district" gorm:"column:sub_district"`
	District      string    `json:"district"`
	Province      string    `json:"province"`
	ZipCode       string    `json:"zip_code" gorm:"column:zip_code"`
}
