package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DogFact struct {
	gorm.Model
	ID		uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	Fact 	string
}