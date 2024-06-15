package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;<-:create"`
	CreatedBy uuid.UUID `gorm:"type:uuid;;<-:create"`
	CreatedAt time.Time `gorm:"autoCreateTime;index;sort:desc;<-:create"`
}
