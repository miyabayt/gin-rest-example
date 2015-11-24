package models

import "github.com/jinzhu/gorm"

type Owner struct {
	gorm.Model
	FirstName string `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
	LastName  string
	Address   string
	City      string
	Telephone string
	Pets      []Pet
}
