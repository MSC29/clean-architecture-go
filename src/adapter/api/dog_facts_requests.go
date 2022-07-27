package api

type DogFactQueryString struct {
	DogFactId   int32 `uri:"dogFactId" binding:"required"`
}