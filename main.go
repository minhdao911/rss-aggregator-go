package main

import (
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}
	fmt.Println("port: ", portString)
}