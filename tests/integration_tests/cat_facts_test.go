package integration_tests

import (
	"clean-architecture/clean-architecture-go/src/adapter/api/cat_facts"
	"clean-architecture/clean-architecture-go/src/infrastructure"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnMultipleCatResults(t *testing.T) {
	// setup
	router := infrastructure.StartApp()
	w := httptest.NewRecorder()

	// given the "all cat facts" route
	req, _ := http.NewRequest("GET", "/api/v1/cats", nil)

	// when getting
	router.ServeHTTP(w, req)

	// then expect entire list
	assert.Equal(t, 200, w.Code)
	response, _ := io.ReadAll(w.Body)

	catPresenters := []cat_facts.CatFactPresenter{}
	jsonErr := json.Unmarshal(response, &catPresenters)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	assert.Equal(t, len(catPresenters), 10)
	assert.Equal(t, catPresenters[0].Fact, "The first true cats came into existence about 12 million years ago and were the Proailurus.")
	assert.Equal(t, catPresenters[0].NbChars, int32(91))
}

func TestReturnOneCatResultOnly(t *testing.T) {
	// setup
	router := infrastructure.StartApp()
	w := httptest.NewRecorder()

	// given the "all cat facts" route
	req, _ := http.NewRequest("GET", "/api/v1/cats/random", nil)

	// when getting
	router.ServeHTTP(w, req)

	// then expect 1 only
	assert.Equal(t, 200, w.Code)
	response, _ := io.ReadAll(w.Body)

	catPresenter := cat_facts.CatFactPresenter{}
	jsonErr := json.Unmarshal(response, &catPresenter)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	assert.Equal(t, catPresenter.Fact, "In the 1930s, two Russian biologists discovered that color change in Siamese kittens depend on their body temperature. Siamese cats carry albino genes that work only when the body temperature is above 98° F. If these kittens are left in a very warm room, their points won’t darken and they will stay a creamy white.")
	assert.Equal(t, catPresenter.NbChars, int32(315))
}
