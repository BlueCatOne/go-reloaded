package main

import (
	"fmt"
	"go-reloaded/src/manipulationFichier"
	"go-reloaded/src/niveaux"
	"go-reloaded/src/tokenizer"
	"os"
)

func main() {
	fichier, err := os.OpenFile("retour.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	manipulationFichier.Verifier(err)
	defer fichier.Close()
	capsule_lue := manipulationFichier.Lire("capsule.txt")
	lignes := tokenizer.Tokenizer(capsule_lue, "\n")
	for _, ligne := range lignes {
		mots := tokenizer.Tokenizer(ligne, " ")
		mots = niveaux.NiveauUn(mots)
		mots = niveaux.NiveauDeux(mots)
		mots = niveaux.NiveauQuatre(mots)
		mots = niveaux.NiveauTrois(mots)
		mots = niveaux.NiveauCinq(mots)
		fmt.Println("------------------------------")
		for i, mot := range mots {
			if i != len(mots)-1 {
				manipulationFichier.Ecrire(mot+" ", fichier)
			}
		}
		manipulationFichier.Ecrire("\n", fichier)
	}
}
