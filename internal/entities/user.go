package entities

type User struct {
	BaseModel
	FullName string `json:"full_name" gorm:"column:full_name"`
	Sex      int    `json:"sex" gorm:"column:sex"`
	Email    string `json:"email" gorm:"column:email;uniqueIndex"`
	Password string `json:"password" gorm:"column:password"`
	IsActive int    `json:"is_active"`
	IsLogin  int    `json:"is_login"`
}
