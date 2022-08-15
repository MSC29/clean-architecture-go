package integration_tests

import (
	"clean-architecture/clean-architecture-go/src/adapter/api/dog_facts"
	"clean-architecture/clean-architecture-go/src/infrastructure"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnMultipleDogResults(t *testing.T) {
	// setup
	router := infrastructure.StartApp()
	w := httptest.NewRecorder()

	// given the "all dog facts" route
	req, _ := http.NewRequest("GET", "/api/v1/dogs", nil)

	// when getting
	router.ServeHTTP(w, req)

	// then expect 3 results (inserted in db)
	assert.Equal(t, 200, w.Code)
	response, _ := io.ReadAll(w.Body)

	dogPresenters := []dog_facts.DogFactPresenter{}
	jsonErr := json.Unmarshal(response, &dogPresenters)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	assert.Equal(t, len(dogPresenters), 3)
	assert.Equal(t, dogPresenters[0].Txt, "Forty-five percent of U.S. dogs sleep in their owner's bed")
	assert.Equal(t, dogPresenters[0].FactId, int32(1))
}

func TestReturnOneDogResultOnly(t *testing.T) {
	// setup
	router := infrastructure.StartApp()
	w := httptest.NewRecorder()

	// given the "single dog facts" route
	dog_fact_id := 2
	req, _ := http.NewRequest("GET", fmt.Sprintf(`/api/v1/dogs/%d`, dog_fact_id), nil)

	// when getting
	router.ServeHTTP(w, req)

	// then expect 1 result (id 2 inserted in db)
	assert.Equal(t, 200, w.Code)
	response, _ := io.ReadAll(w.Body)

	dogPresenter := dog_facts.DogFactPresenter{}
	jsonErr := json.Unmarshal(response, &dogPresenter)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	assert.Equal(t, dogPresenter.Txt, "Seventy percent of people sign their dog's name on their holiday cards")
	assert.Equal(t, dogPresenter.FactId, int32(2))
}
