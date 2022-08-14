package main

import (
	"clean-architecture/clean-architecture-go/src/infrastructure"

	"fmt"
)

func main() {
	fmt.Println("app starting")
	router := infrastructure.StartApp()

	router.Run("localhost:8080")
}
