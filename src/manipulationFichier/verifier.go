package manipulationFichier

import (
	"fmt"
	"os"
)

func Verifier(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}