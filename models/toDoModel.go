package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
