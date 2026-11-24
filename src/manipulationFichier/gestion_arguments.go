package manipulationFichier

import (
	"fmt"
	"os"
)

func Gestion_arguments() (string, string) {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("arguments invalides")
		os.Exit(1)
	}
	entree := args[1]
	sortie := args[2]
	return entree, sortie
}
