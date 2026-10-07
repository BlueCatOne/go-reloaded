package tokenazer

func Tokenizer(phrase, separateur string) []string {
	var slicefinale []string
	var mot string
	space := len(separateur) // int
	for indice := 0; indice < len(phrase); indice++ {
		if !(indice+space < len(phrase)-1 && separateur == string(phrase[indice:indice+space])) {
			if indice == len(phrase)-1 {
				mot += string(phrase[indice])
				if mot != "" {
					slicefinale = append(slicefinale, mot)
				}
			} else {
				mot += string(phrase[indice])
			}
		} else {
			if mot != "" {
				slicefinale = append(slicefinale, mot)
			}
			indice += space - 1
			mot = ""
		}
	}
	return slicefinale
}
