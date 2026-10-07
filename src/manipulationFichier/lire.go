package manipulationFichier

import "os"

func Lire(filename string) string {
	data, err := os.ReadFile(filename)
	Verifier(err)
	return string(data)
}
