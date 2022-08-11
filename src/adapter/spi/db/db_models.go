package db

import (
	"gorm.io/gorm"
)

type DogFact struct {
	gorm.Model
	ID   int32 `gorm:"primaryKey;autoIncrement:true"`
	Fact string
}
