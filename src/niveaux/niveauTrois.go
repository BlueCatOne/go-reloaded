package niveaux

//i=position du mot ; ligne[i]=mot ; ligne[i][0] = 1ère lettre du mot
func NiveauTrois(ligne []string) []string {
	ligne_finale := []string{ligne[0]}
	for i := 1; i < len(ligne); i++ {
		if len(ligne[i]) != 0 && !(ligne[i][0] == ',' || ligne[i][0] == '.' || ligne[i][0] == '!' || ligne[i][0] == '?' || ligne[i][0] == ':' || ligne[i][0] == ';') {
			ligne_finale = append(ligne_finale, ligne[i])
		} else {
			for len(ligne[i]) != 0 && (ligne[i][0] == ',' || ligne[i][0] == '.' || ligne[i][0] == '!' || ligne[i][0] == '?' || ligne[i][0] == ':' || ligne[i][0] == ';') {
				ligne_finale[len(ligne_finale)-1] += string(ligne[i][0]) // concaténation de la 1ère lettre du mot (ligne[i][0]) avec le mot précédent (ligne[i-1])
				ligne[i] = ligne[i][1:]                                  // supprime la première lettre du mot en ne gardant que ce qui suit à partir de l'index 1
				i -= 1                                                   // revérification du nouveau mot actuel
			}
		}
	}
	return ligne_finale
}
