package cat_facts

import (
	"net/http"

	"clean-architecture/clean-architecture-go/src/application/repositories"
	"clean-architecture/clean-architecture-go/src/application/usecases"

	"github.com/gin-gonic/gin"
)

type CatFactsController struct {
    Router *gin.Engine
	CatFactsRepository repositories.CatFactsRepositoryAbstract
}

func (ctrl *CatFactsController) SetupRoutes(){
	ctrl.Router.GET("/api/v1/cats/", ctrl.GetAllCatFacts)
	ctrl.Router.GET("/api/v1/cats/random", ctrl.GetOneRandomCatFactById)
}

func (ctrl *CatFactsController) GetAllCatFacts(c *gin.Context) {
	catFactPresenterMapper := CatFactPresenterMapper{}

	useCase := usecases.GetAllCatFactsUseCase{ Repository: ctrl.CatFactsRepository }
	catFacts := useCase.Execute()

	var catFactsPresenter []CatFactPresenter
	var i int
	for i = 0; i < len(catFacts); i++ {
		catFactsPresenter = append(catFactsPresenter, catFactPresenterMapper.toApi(catFacts[i]))
	}
	
    c.IndentedJSON(http.StatusOK, catFactsPresenter)
}

func (ctrl *CatFactsController) GetOneRandomCatFactById(c *gin.Context) {    
	catFactPresenterMapper := CatFactPresenterMapper{}

	useCase := usecases.GetOneRandomCatFactUseCase{ Repository: ctrl.CatFactsRepository }
	catFact := useCase.Execute()

	c.IndentedJSON(http.StatusOK, catFactPresenterMapper.toApi(catFact))
}