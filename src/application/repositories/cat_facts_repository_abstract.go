package repositories

import "clean-architecture/clean-architecture-go/src/domain"

type CatFactsRepositoryAbstract interface {
	GetRandomCatFact() domain.CatFactEntity
	GetAllCatFacts() []domain.CatFactEntity
}