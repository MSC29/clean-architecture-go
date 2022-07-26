package domain

import (
	"time"
)

type DogFactEntity struct {
	FactId int32
	Fact   string
}

type CatFactEntity struct {
	FactTxt    string
	FactLength int32
}

type ApiError struct {
	When time.Time
	What string
}

type ConfigurationEntity struct {
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	CatsSource       string
	Env              string
}
