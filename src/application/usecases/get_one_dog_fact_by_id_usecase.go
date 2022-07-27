package usecases

import (
	"clean-architecture/clean-architecture-go/src/application/repositories"
	"clean-architecture/clean-architecture-go/src/domain"
)

type GetOneDogFactByIdUseCase struct {
	DogFactId int32
    Repository repositories.DogFactsRepositoryAbstract
}

func (useCase *GetOneDogFactByIdUseCase) Execute() domain.DogFactEntity {
    return useCase.Repository.GetDogFactById(useCase.DogFactId)
}