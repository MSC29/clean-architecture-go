package db

import (
	"clean-architecture/clean-architecture-go/src/domain"
)
  
type DogFactsRepository struct{
	DbConnection DbConnection
	DogFactDbMapper DogFactDbMapper
}

func NewDogFactsRepository(dbConnection DbConnection) *DogFactsRepository{
	return &DogFactsRepository{DbConnection: dbConnection, DogFactDbMapper: DogFactDbMapper{}}
}

func (repo *DogFactsRepository) GetDogFactById(factId int32) domain.DogFactEntity{
	var dogFact DogFact
	repo.DbConnection.GetManager().Find(&dogFact, factId)
	
	return repo.DogFactDbMapper.toEntity(dogFact)
}

func (repo *DogFactsRepository) GetAllDogFacts() []domain.DogFactEntity{
	var dogFacts []DogFact
	repo.DbConnection.GetManager().Find(&dogFacts)

	var resultEntities []domain.DogFactEntity
	for i := 0; i < len(dogFacts); i++ {
		resultEntities = append(resultEntities, repo.DogFactDbMapper.toEntity(dogFacts[i]))
	}

	return resultEntities
}
