package api

import (
	"clean-architecture/clean-architecture-go/src/domain"
)

type DogFactPresenterMapper struct{}

func (mapper *DogFactPresenterMapper) toApi(entity domain.DogFactEntity) DogFactPresenter {
		return DogFactPresenter{
			FactId: entity.FactId,
			Txt: entity.Fact,
		}
	}

// func (mapper *DogFactPresenterMapper) toEntity(payload BasePayloadEntity) domain.DogFactEntity {
// 	return domain.ApiError{time.Now(), "not implemented"}
// }
