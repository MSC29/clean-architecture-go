package main

import (
	"clean-architecture/clean-architecture-go/src/infrastructure"

	"fmt"
)

func main() {
	fmt.Println("app starting")
	infrastructure.StartApp()
}