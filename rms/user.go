package rms

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	IsDeleted    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ID           uint
	Name         string
	Phone        string
	Email        string
	Password     string
	Department   string
	Role         string
	IsSuperAdmin int
}

func GetAllUserBySearch(db *gorm.DB, searchUser User) []User {
	var defaultUser User
	var users []User
	if searchUser == defaultUser {
		db.Table("user").Find(&users)
	} else {
		db.Table("user").Find(&users, searchUser)
	}

	return users
}
