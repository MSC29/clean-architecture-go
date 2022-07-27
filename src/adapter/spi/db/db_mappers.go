package db

import "clean-architecture/clean-architecture-go/src/domain"

type DogFactDbMapper struct {}

func (mapper *DogFactDbMapper) toEntity(model DogFact) domain.DogFactEntity {
	entity := domain.DogFactEntity{
		FactId: model.ID,
		Fact: model.Fact,
	}

	return entity;
}

func (mapper *DogFactDbMapper) toModel(entity domain.DogFactEntity) DogFact {
	model := DogFact{
		ID: entity.FactId,
		Fact: entity.Fact,
	}

	return model;
}
