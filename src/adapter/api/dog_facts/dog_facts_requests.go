package dog_facts

type DogFactQueryString struct {
	DogFactId   int32 `uri:"dogFactId" binding:"required"`
}