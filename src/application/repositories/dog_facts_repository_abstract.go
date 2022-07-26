package repositories

type DogFactsRepositoryAbstract interface {
	GetDogFactById(factId int)
	GetAllDogFacts()
}