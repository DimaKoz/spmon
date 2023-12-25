package main

import (
	"log"

	"github.com/DimaKoz/spmon/internal/repository"
)

func main() {
	log.Println("Hello world")

	a, err := repository.GetHs("")
	log.Println(err, a)
}
