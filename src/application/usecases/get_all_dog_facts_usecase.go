package usecases

import (
	"fmt"

	"clean-architecture/clean-architecture-go/src/application/repositories"
)

type GetAllDogFactsUseCase struct {
    Repository repositories.DogFactsRepositoryAbstract
}

func (useCase *GetAllDogFactsUseCase) Execute() {
    fmt.Println("GetAllDogFactsUseCase")
	useCase.Repository.GetAllDogFacts()
}