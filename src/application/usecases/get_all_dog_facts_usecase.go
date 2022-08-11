package usecases

import (
	"clean-architecture/clean-architecture-go/src/application/repositories"
	"clean-architecture/clean-architecture-go/src/domain"
)

type GetAllDogFactsUseCase struct {
	Repository repositories.DogFactsRepositoryAbstract
}

func (useCase *GetAllDogFactsUseCase) Execute() []domain.DogFactEntity {
	return useCase.Repository.GetAllDogFacts()
}
