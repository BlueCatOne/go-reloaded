package niveaux

func GuillemetFutur(ligne []string, départ int) bool {
	for i := départ; i < len(ligne); i++ {
		if ligne[i] == "'" || (len(ligne[i]) > 0 && string(ligne[i][len(ligne[i])-1]) == "'") {
			return true
		}
	}
	return false
}

// i=position du mot ; ligne = phrase ; ligne[i]=mot ; ligne[i][0] = 1ère lettre du mot ; ligne_finale[len(ligne_finale)-1] = dernier mot de ligne_finale
func NiveauQuatre(ligne []string) []string {
	ligne_finale := []string{}
	ouvrant := true
	for i := 0; i < len(ligne); i++ {
		if ligne[i] == "'" {
			if ouvrant == true {
				if i != len(ligne)-1 && GuillemetFutur(ligne, i+1) {
					ligne[i+1] = "'" + ligne[i+1]
				} else {
					ligne_finale = append(ligne_finale, ligne[i])
				}
			} else {
				ouvrant = true
				ligne_finale[len(ligne_finale)-1] += "'"
			}
			continue
		}
		if string(ligne[i][0]) == "'" { // si il y a un ' au début du mot
			if ouvrant == true {
				if !GuillemetFutur(ligne, i+1) {
					ligne_finale = append(ligne_finale, ligne[i])
					continue
				}
				ouvrant = false
			} else {
				ligne_finale[len(ligne_finale)-1] += "'"
				ouvrant = true
				ligne[i] = ligne[i][1:]
			}
		}
		if string(ligne[i][len(ligne[i])-1]) == "'" { // si il y a un ' à la fin du mot
			if ouvrant == true && i != len(ligne)-1 {
				ligne[i] = ligne[i][:len(ligne[i])-1]
				ligne[i+1] = "'" + ligne[i+1]
			} else {
				ouvrant = !ouvrant // si ouvrant = true --> false, si ouvrant = false --> true
			}
		}
		ligne_finale = append(ligne_finale, ligne[i])
	}
	return ligne_finale
}
