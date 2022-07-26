package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
type DbConnection struct{
	Db *gorm.DB
}

func NewDbConnection() DbConnection{
	dsn := "host=localhost user=postgres password=postgres dbname=animal_fact_db_go port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&DogFact{})

	return DbConnection{Db: db}
}

func (dbConnection *DbConnection) GetManager() *gorm.DB{
	return dbConnection.Db
}