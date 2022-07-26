package infrastructure

import (
	"clean-architecture/clean-architecture-go/src/adapter/api"
	"clean-architecture/clean-architecture-go/src/adapter/spi/db"
	"clean-architecture/clean-architecture-go/src/application/repositories"

	"github.com/gin-gonic/gin"
)

func StartApp() {
    router := gin.Default()

	dbConnection := db.NewDbConnection()

	var dogFactsRepository repositories.DogFactsRepositoryAbstract
	dogFactsRepository = db.NewDogFactsRepository(dbConnection)

	pCtrl := api.DogFactsController{ Router: router, DogFactsRepository: dogFactsRepository }
	pCtrl.SetupRoutes()

    router.Run("localhost:8080")
}