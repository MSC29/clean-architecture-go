package usecases

import (
	"clean-architecture/clean-architecture-go/src/application/repositories"
	"clean-architecture/clean-architecture-go/src/domain"
)

type GetOneRandomCatFactUseCase struct {
	Repository repositories.CatFactsRepositoryAbstract
}

func (useCase *GetOneRandomCatFactUseCase) Execute() domain.CatFactEntity {
	return useCase.Repository.GetRandomCatFact()
}
