package integration_tests

import (
	"clean-architecture/clean-architecture-go/src/adapter/api/dog_facts"
	"clean-architecture/clean-architecture-go/src/infrastructure"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")

	// Return a function to teardown the test
	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

// Almost the same as the above, but this one is for single test instead of collection of tests
func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	// app := infrastructure.StartApp()
	return func(tb testing.TB) {
		log.Println("teardown test")
	}
}

func TestReturnMultipleResults(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	router := infrastructure.StartApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/dogs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())

	response, _ := ioutil.ReadAll(w.Body)

	dogPresenters := []dog_facts.DogFactPresenter{}
	jsonErr := json.Unmarshal(response, &dogPresenters)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	assert.Equal(t, len(dogPresenters), 3)
	assert.Equal(t, dogPresenters[0].Txt, "Forty-five percent of U.S. dogs sleep in their owner's bed")
	assert.Equal(t, dogPresenters[0].FactId, int32(1))
}
