package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;<-:create"`
	FullName  string    `gorm:"column:full_name"`
	Sex       int       `gorm:"column:sex"`
	Email     string    `gorm:"column:email;uniqueIndex"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	IsActive  int
	IsLogin   int
	CreatedAt time.Time `gorm:"autoCreateTime;index;sort:desc;<-:create"`
}

func (User) TableName() string {
	return "user"
}
