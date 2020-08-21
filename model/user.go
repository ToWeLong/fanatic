package model

import (
	"time"
)

type User struct {
	ID        uint `json:"id" gorm:"primary_key"`
	Account string `json:"account" gorm:"unique;not null"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
