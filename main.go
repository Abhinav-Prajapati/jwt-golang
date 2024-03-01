package main

import (
	"fmt"
	"go-jwt/initializers"
)

func main() {
	fmt.Println("hello")
	fmt.Println("test")
}

func init() {
	initializers.LoadEvnVariables()
}
