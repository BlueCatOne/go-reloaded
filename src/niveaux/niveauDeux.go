package niveaux

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func capitalize(s string) string {
	capitalizing := true
	subtext := ""
	toreturn := ""
	for _, lettre := range s {

		subtext = string(lettre)

		if isLower(subtext) && capitalizing {
			subtext = toUpper(subtext)
		}

		if !capitalizing {
			subtext = toLower(subtext)
		}

		if isAlpha(subtext) {
			capitalizing = false
		} else {
			capitalizing = true
		}
		toreturn += subtext
	}
	return toreturn
}

func isAlpha(s string) bool {
	for i, l := range s {
		if !((l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (l >= '0' && l <= '9')) {
			i = i
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, l := range s {
		if !(l >= 'a' && l <= 'z') {
			return false
		}
	}
	return true
}

func toUpper(s string) string {
	mot := []rune(s)
	for i := range mot {
		if mot[i] >= 'a' && mot[i] <= 'z' {
			mot[i] -= 32
		}
	}
	return string(mot)
}

func toLower(s string) string {
	mot := []rune(s)
	for i := range mot {
		if mot[i] >= 'A' && mot[i] <= 'Z' {
			mot[i] += 32
		}
	}
	return string(mot)
}

func NiveauDeux(ligne []string) []string {
	nouvelle_ligne := []string{ligne[0]}
	for i := 1; i < len(ligne); i++ {
		//cas sans nombres
		if ligne[i] == "(up)" {
			nouvelle_ligne = appliquer_modifieur(nouvelle_ligne, toUpper, 1)
		} else if ligne[i] == "(low)" {
			nouvelle_ligne = appliquer_modifieur(nouvelle_ligne, toLower, 1)
		} else if ligne[i] == "(cap)" {
			nouvelle_ligne = appliquer_modifieur(nouvelle_ligne, capitalize, 1)
			
		//cas avec nombres
		} else if len(ligne[i]) >= 6 && ligne[i][:5] == "(up, " && ligne[i][len(ligne[i])-1:] == ")" { // up

			nb_mod, err := strconv.Atoi(ligne[i][5 : len(ligne[i])-1])
			if err != nil {
				fmt.Println("Erreur de conversion:", err)
				os.Exit(1)
			}
			nouvelle_ligne = appliquer_modifieur(nouvelle_ligne, toUpper, nb_mod)

		} else if len(ligne[i]) >= 7 && ligne[i][:6] == "(low, " && ligne[i][len(ligne[i])-1:] == ")" { // low

			nb_mod, err := strconv.Atoi(ligne[i][6 : len(ligne[i])-1])
			if err != nil {
				fmt.Println("Erreur de conversion:", err)
				os.Exit(1)
			}
			nouvelle_ligne = appliquer_modifieur(nouvelle_ligne, toLower, nb_mod)

		} else if len(ligne[i]) >= 7 && ligne[i][:6] == "(cap, " && ligne[i][len(ligne[i])-1:] == ")" { // cap

			nb_mod, err := strconv.Atoi(ligne[i][6 : len(ligne[i])-1])
			if err != nil {
				fmt.Println("Erreur de conversion:", err)
				os.Exit(1)
			}
			nouvelle_ligne = appliquer_modifieur(nouvelle_ligne, capitalize, nb_mod)
		} else {
			nouvelle_ligne = append(nouvelle_ligne, ligne[i])
		}
	}
	return nouvelle_ligne
}

func appliquer_modifieur(nouvelle_ligne []string, f func(string) string, nb int) []string {
	tomodify:= []string{}
	fmt.Println("nb_mod:", nb)
	for j := 0; j < nb; j++ {
		tomodify = append(tomodify, nouvelle_ligne[len(nouvelle_ligne)-1])
		if !isAlpha(nouvelle_ligne[len(nouvelle_ligne)-1]) {
			j--
		}
		nouvelle_ligne = nouvelle_ligne[:len(nouvelle_ligne)-1]
	}
	slices.Reverse(tomodify)
	for _,mot := range tomodify {
		nouvelle_ligne = append(nouvelle_ligne, f(mot))
	}
	return nouvelle_ligne
}