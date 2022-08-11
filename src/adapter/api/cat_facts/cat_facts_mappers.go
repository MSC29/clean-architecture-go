package cat_facts

import (
	"clean-architecture/clean-architecture-go/src/domain"
)

type CatFactPresenterMapper struct{}

func (mapper *CatFactPresenterMapper) toApi(entity domain.CatFactEntity) CatFactPresenter {
	return CatFactPresenter{
		Fact:    entity.FactTxt,
		NbChars: entity.FactLength,
	}
}

// func (mapper *CatFactPresenterMapper) toEntity(payload BasePayloadEntity) domain.CatFactEntity {
// 	return domain.ApiError{time.Now(), "not implemented"}
// }
