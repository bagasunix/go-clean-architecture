package entities

import "github.com/gofrs/uuid"

type AddressCustomer struct {
	BaseModel
	CustomerID    uuid.UUID
	Customer      *Customer `gorm:"foreignKey:CustomerID"`
	StreetAddress string    `gorm:"column:street_address"`
	SubDistrict   string    `gorm:"column:sub_district"`
	District      string
	Province      string
	ZipCode       string `gorm:"column:zip_code"`
}
