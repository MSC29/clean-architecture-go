package usecases

import (
	"clean-architecture/clean-architecture-go/src/application/repositories"
	"clean-architecture/clean-architecture-go/src/domain"
)

type GetAllCatFactsUseCase struct {
	Repository repositories.CatFactsRepositoryAbstract
}

func (useCase *GetAllCatFactsUseCase) Execute() []domain.CatFactEntity {
	return useCase.Repository.GetAllCatFacts()
}
