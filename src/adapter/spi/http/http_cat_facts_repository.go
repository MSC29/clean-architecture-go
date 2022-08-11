package http

import (
	"clean-architecture/clean-architecture-go/src/domain"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type CatFactsRepository struct {
	HttpConnection    HttpConnection
	CatFactHttpMapper CatFactHttpMapper
	Source            string
}

func NewCatFactsRepository(httpConnection HttpConnection, source string) *CatFactsRepository {
	return &CatFactsRepository{HttpConnection: httpConnection, Source: source, CatFactHttpMapper: CatFactHttpMapper{}}
}

func (repo *CatFactsRepository) GetRandomCatFact() domain.CatFactEntity {
	response, err := repo.HttpConnection.Client.Get(fmt.Sprintf(`%s/fact`, repo.Source))
	if err != nil {
		log.Fatal(err)
	}

	httpData, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	catFact := CatFactApiModel{}
	jsonErr := json.Unmarshal(httpData, &catFact)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return repo.CatFactHttpMapper.toEntity(catFact)
}

func (repo *CatFactsRepository) GetAllCatFacts() []domain.CatFactEntity {
	response, err := repo.HttpConnection.Client.Get(fmt.Sprintf(`%s/facts`, repo.Source))
	if err != nil {
		log.Fatal(err)
	}

	httpData, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var catFacts CatFactsApiModel
	jsonErr := json.Unmarshal(httpData, &catFacts)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var resultEntities []domain.CatFactEntity
	for i := 0; i < len(catFacts.Data); i++ {
		resultEntities = append(resultEntities, repo.CatFactHttpMapper.toEntity(catFacts.Data[i]))
	}

	return resultEntities
}
