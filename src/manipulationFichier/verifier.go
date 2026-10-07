package manipulationFichier

import "os"

func Verifier(e error) {
	if e != nil {
		os.Exit(1)
	}
}