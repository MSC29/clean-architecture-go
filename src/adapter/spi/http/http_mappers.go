package http

import "clean-architecture/clean-architecture-go/src/domain"

type CatFactHttpMapper struct {}

func (mapper *CatFactHttpMapper) toEntity(model CatFactApiModel) domain.CatFactEntity {
	entity := domain.CatFactEntity{
		FactTxt: model.Fact,
		FactLength: model.Length,
	}

	return entity;
}

func (mapper *CatFactHttpMapper) toModel(entity domain.CatFactEntity) CatFactApiModel {
	model := CatFactApiModel{
		Length: entity.FactLength,
		Fact: entity.FactTxt,
	}

	return model;
}
