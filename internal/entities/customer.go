package entities

type Customer struct {
	BaseModel
	FullName string            `gorm:"column:full_name"`
	Sex      int               `gorm:"column:sex"`
	Phone    string            `gorm:"column:phone;uniqueIndex"`
	Email    string            `gorm:"uniqueIndex"`
	UserName string            `gorm:"column:username"`
	Password string            `gorm:"column:password"`
	IsActive int               `gorm:"column:is_active"`
	IsALogin int               `gorm:"column:is_login"`
	Address  []AddressCustomer `gorm:"foreignKey:CustomerID"`
}
