package dog_facts

import (
	"net/http"

	"clean-architecture/clean-architecture-go/src/application/repositories"
	"clean-architecture/clean-architecture-go/src/application/usecases"

	"github.com/gin-gonic/gin"
)

type DogFactsController struct {
    Router *gin.Engine
	DogFactsRepository repositories.DogFactsRepositoryAbstract
}

func (ctrl *DogFactsController) SetupRoutes(){
	ctrl.Router.GET("/api/v1/dogs/", ctrl.GetAllDogFacts)
	ctrl.Router.GET("/api/v1/dogs/:dogFactId", ctrl.GetOneDogFactById)
}

func (ctrl *DogFactsController) GetAllDogFacts(c *gin.Context) {
	dogFactPresenterMapper := DogFactPresenterMapper{}

	useCase := usecases.GetAllDogFactsUseCase{ Repository: ctrl.DogFactsRepository }
	dogFacts := useCase.Execute()

	var dogFactsPresenter []DogFactPresenter
	var i int
	for i = 0; i < len(dogFacts); i++ {
		dogFactsPresenter = append(dogFactsPresenter, dogFactPresenterMapper.toApi(dogFacts[i]))
	}
	
    c.IndentedJSON(http.StatusOK, dogFactsPresenter)
}

func (ctrl *DogFactsController) GetOneDogFactById(c *gin.Context) {
	var dogFactQueryString DogFactQueryString
	c.ShouldBindUri(&dogFactQueryString)
    
	dogFactPresenterMapper := DogFactPresenterMapper{}

	useCase := usecases.GetOneDogFactByIdUseCase{ DogFactId: dogFactQueryString.DogFactId, Repository: ctrl.DogFactsRepository }
	dogFact := useCase.Execute()

	c.IndentedJSON(http.StatusOK, dogFactPresenterMapper.toApi(dogFact))
}