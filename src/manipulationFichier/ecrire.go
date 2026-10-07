package manipulationFichier

import "os"

func Ecrire(text string, file *os.File) {
    if _, err := file.WriteString(text); err != nil {
        panic(err)
    }
}