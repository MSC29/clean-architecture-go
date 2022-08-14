package db

import (
	"clean-architecture/clean-architecture-go/src/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Db *gorm.DB
}

func NewDbConnection(configuration domain.ConfigurationEntity) DbConnection {
	dsn := fmt.Sprintf(`host=localhost user=%s password=%s dbname=%s port=5432`,
		configuration.DatabaseUsername,
		configuration.DatabasePassword,
		configuration.DatabaseName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&DogFact{})

	return DbConnection{Db: db}
}

func (dbConnection *DbConnection) GetManager() *gorm.DB {
	return dbConnection.Db
}
