package domain

import (
	"time"
)

type DogFactEntity struct {
	FactId int32
	Fact string
}

type ApiError struct {
	When time.Time
	What string
}