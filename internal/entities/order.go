package entities

import "github.com/gofrs/uuid"

type Order struct {
	BaseModel
	CustomerID   uuid.UUID `gorm:"column:customer_id"`
	AddressID    uuid.UUID `gorm:"column:address_id"`
	ProductName  string    `gorm:"column:product_name"`
	Amount       int
	ProductPrice float64 `gorm:"column:product_price"`
	Status       string
}
