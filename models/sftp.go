package models

import "gorm.io/gorm"

type SFTPConfiguration struct {
	Username string
	Password string
	Host     string
	Port     string
	BasePath string
}

type CSVRawFile struct {
	FileName string
	RawFile  []byte
}

type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
}

type GormPerson struct {
	gorm.Model
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	Email       string `gorm:"column:email"`
	PhoneNumber string `gorm:"column:phone_number"`
	DateOfBirth string `gorm:"column:date_of_birth"`
	Address     string `gorm:"column:address"`
}
