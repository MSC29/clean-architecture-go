package infrastructure

import (
	"clean-architecture/clean-architecture-go/src/adapter/api/cat_facts"
	"clean-architecture/clean-architecture-go/src/adapter/api/dog_facts"
	"clean-architecture/clean-architecture-go/src/adapter/spi/db"
	"clean-architecture/clean-architecture-go/src/adapter/spi/http"
	"clean-architecture/clean-architecture-go/src/application/repositories"

	"github.com/gin-gonic/gin"
)

func StartApp() {
    router := gin.Default()

	dbConnection := db.NewDbConnection()
	httpConnection := http.NewHttpConnection()

	var dogFactsRepository repositories.DogFactsRepositoryAbstract
	dogFactsRepository = db.NewDogFactsRepository(dbConnection)

	var catFactsRepository repositories.CatFactsRepositoryAbstract
	catFactsRepository = http.NewCatFactsRepository(httpConnection, "https://catfact.ninja")

	dogController := dog_facts.DogFactsController{ Router: router, DogFactsRepository: dogFactsRepository }
	dogController.SetupRoutes()
	catController := cat_facts.CatFactsController{ Router: router, CatFactsRepository: catFactsRepository }
	catController.SetupRoutes()

    router.Run("localhost:8080")
}