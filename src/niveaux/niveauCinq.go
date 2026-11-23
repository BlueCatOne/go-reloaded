package niveaux

func NiveauCinq(ligne []string) []string {
	for i := 0; i < len(ligne)-1; i++ {
		if ligne[i] == "a" || ligne[i] == "an" {
			compteur := 0
			for ligne[i+1][compteur] == '\'' && compteur < len(ligne[i+1]) { // compte le nombre d'apostrophes au début du mot
				compteur++
			}
			if ligne[i+1][compteur] == 'a' || ligne[i+1][compteur] == 'e' || ligne[i+1][compteur] == 'i' || ligne[i+1][compteur] == 'o' || ligne[i+1][compteur] == 'u' {
				ligne[i] = "an"
			} else if ligne[i+1][compteur] >= 'a' && ligne[i+1][compteur] <= 'z' {
				ligne[i] = "a"
			}
		} else if ligne[i] == "A" || ligne[i] == "An" {
			compteur := 0
			for ligne[i+1][compteur] == '\'' && compteur < len(ligne[i+1]) { // compte le nombre d'apostrophes au début du mot
				compteur++
			}
			if ligne[i+1][compteur] == 'a' || ligne[i+1][compteur] == 'e' || ligne[i+1][compteur] == 'i' || ligne[i+1][compteur] == 'o' || ligne[i+1][compteur] == 'u' {
				ligne[i] = "An"
			} else if ligne[i+1][compteur] >= 'a' && ligne[i+1][compteur] <= 'z' {
				ligne[i] = "A"
			}
		}
	}
	return ligne
}
