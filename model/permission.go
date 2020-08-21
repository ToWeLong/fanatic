package model

import "time"

type Permission struct {
	ID        uint `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"unique;not null"`
	Module string `json:"module"`
	Mount int `json:"mount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
