package entity

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string  `gorm:"<-:create" json:"name"`
	Email        *string `gorm:"index" json:"email"`
	Age          uint8   `json:"age"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}
