package repositories

import "clean-architecture/clean-architecture-go/src/domain"

type DogFactsRepositoryAbstract interface {
	GetDogFactById(factId int32) domain.DogFactEntity
	GetAllDogFacts() []domain.DogFactEntity
}