package arguments

import (
	"fmt"
	"os"
)

func GestionArguments(args []string) (string, string) {
	switch {
	case len(args) < 3 || len(args) > 4:
		fmt.Println("arguments invalides !")
		os.Exit(1)
	case len(args) == 4 && args[3] == "--smart":
		return args[1], args[2]
	case len(args) == 3:
		return args[1], args[2]
	default:
		fmt.Println("arguments invalides !")
		os.Exit(1)
	}
	return args[1], args[2]
}
