package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Pet struct {
	gorm.Model
	Birthday time.Time `sql:"DEFAULT:current_timestamp"`
	Name     string
	Age      int
	OwnerId  int64 `sql:"index"`
}
