package model

import (
	"github.com/template/go-backend-gin-orm/dtos"
)

type User struct {
	dtos.BaseModel
	Name  string // A regular string field
	Email string `gorm:"unique"` // A pointer to a string, allowing for null values
	Bio   *string
	// MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	MemberNumber *string // Uses sql.NullString to handle nullable strings
	Avatar       *string

	Books []Book `gorm:"foreignKey:UserID"`
	Login *Login `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type Login struct {
	dtos.BaseModel
	Email        string `gorm:"uniqueIndex"`
	UserID       uint   `gorm:"unique"`
	HashPassword string

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Book struct {
	dtos.BaseModel
	UserID    uint
	Title     string
	Author    string
	Publisher string

	User User `gorm:"constraint:OnDelete:CASCADE"`
}

func GetAllModels() []interface{} {
	return []interface{}{
		&User{},
		&Book{},
		&Login{},
	}
}
