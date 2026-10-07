package main

import (
	"go-reloaded/src/manipulationFichier"
	"os"
)

func main() {
	file, err := os.OpenFile("test2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	manipulationFichier.Verifier(err)
	defer file.Close()
	
}
