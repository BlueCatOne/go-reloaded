package niveaux

import (
	"fmt"
	"slices"
)

// fonctions utilitaires pour le fonctionnement du Niveau 1
func indice(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func iterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	total := 1
	for i := 0; i < power; i++ {
		total *= nb
	}
	return total
}

func bintodec(nombre string) string {
	if nombre == "" {
		return ""
	}
	carac := []string{"0", "1"}
	resultat := 0
	for i, chiffre := range nombre {
		resultat += iterativePower(len(carac), len(nombre)-1-i) * indice(string(chiffre), carac)
		if !slices.Contains(carac, string(chiffre)) {
			return nombre
		}
	}
	return fmt.Sprint(resultat)
}

func hextodec(nombre string) string {
	if nombre == "" {
		return ""
	}
	carac := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	resultat := 0
	for i, chiffre := range nombre {
		resultat += iterativePower(len(carac), len(nombre)-1-i) * indice(string(chiffre), carac)
		if !slices.Contains(carac, string(chiffre)) {
			return nombre
		}
	}
	return fmt.Sprint(resultat)
}

// fin des fonctions utilitaires.

func NiveauUn(ligne []string) []string {
	nouvelle_ligne := []string{}
	for i := 0; i < len(ligne); i++ {
		if ligne[i] == "(hex)" && i != 0{
			nouvelle_ligne = nouvelle_ligne[:len(nouvelle_ligne)-1]
			nouvelle_ligne = append(nouvelle_ligne, fmt.Sprint(hextodec(ligne[i-1])))
		} else if ligne[i] == "(bin)" && i != 0{
			nouvelle_ligne = nouvelle_ligne[:len(nouvelle_ligne)-1]
			nouvelle_ligne = append(nouvelle_ligne, fmt.Sprint(bintodec(ligne[i-1])))
		} else {
			nouvelle_ligne = append(nouvelle_ligne, ligne[i])
		}
	}
	return nouvelle_ligne
}
