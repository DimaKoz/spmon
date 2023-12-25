package main

import (
	"github.com/DimaKoz/spmon/internal/repository"
	"log"
)

func main() {
	log.Println("Hello world")

	a, err := repository.GetHs("")
	log.Println(err, a)
}
