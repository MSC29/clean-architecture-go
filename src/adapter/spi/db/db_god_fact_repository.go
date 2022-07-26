package db

import (
	"fmt"
)
  
type DogFactsRepository struct{
	DbConnection DbConnection
}

func NewDogFactsRepository(dbConnection DbConnection) *DogFactsRepository{
	return &DogFactsRepository{DbConnection: dbConnection}
}

func (repo *DogFactsRepository) GetDogFactById(factId int){
	fmt.Printf("getDogFactById: %T", factId)
}

func (repo *DogFactsRepository) GetAllDogFacts(){
	fmt.Printf("getAllDogFacts")
	
	var dogFacts []DogFact
	result := repo.DbConnection.GetManager().Find(&dogFacts)

	fmt.Printf("GetAllDogFacts: RowsAffected %T, Error %T", result.RowsAffected, result.Error)
}
