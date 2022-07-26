package api

import (
	"fmt"
	"net/http"

	"clean-architecture/clean-architecture-go/src/application/repositories"
	"clean-architecture/clean-architecture-go/src/application/usecases"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type DogFactsController struct {
    Router *gin.Engine
	DogFactsRepository repositories.DogFactsRepositoryAbstract
}

func init() {
	fmt.Println("init DogFactsController")
}

func (ctrl *DogFactsController) SetupRoutes(){
	ctrl.Router.GET("/api/v1/dogs/", ctrl.GetAllDogFacts)
	ctrl.Router.GET("/api/v1/dogs/:dogFactId", ctrl.GetOneDogFactById)
}

func (ctrl *DogFactsController) GetAllDogFacts(c *gin.Context) {
	useCase := usecases.GetAllDogFactsUseCase{ Repository: ctrl.DogFactsRepository }
	useCase.Execute()
    c.IndentedJSON(http.StatusOK, albums)
}

func (ctrl *DogFactsController) GetOneDogFactById(c *gin.Context) {
	var dogFactQueryString DogFactQueryString
	c.ShouldBindUri(&dogFactQueryString)
    c.IndentedJSON(http.StatusOK, dogFactQueryString.ID)
}