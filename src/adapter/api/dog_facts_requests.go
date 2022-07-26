package api

type DogFactQueryString struct {
	ID   string `uri:"dogFactId" binding:"required,uuid"`
}