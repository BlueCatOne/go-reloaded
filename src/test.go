package main

import (
	"fmt"
	"go-reloaded/src/niveaux"
	"go-reloaded/src/tokenizer"
)

func main() {
	fmt.Println(niveaux.NiveauTrois(tokenizer.Tokenizer("I was sitting over there ,and then BAMM !!", " ")))
	fmt.Println(niveaux.NiveauTrois(tokenizer.Tokenizer("I was thinking ... You were right", " ")))
	fmt.Println(niveaux.NiveauTrois(tokenizer.Tokenizer("Really ! ? No way", " ")))
	fmt.Println(niveaux.NiveauTrois(tokenizer.Tokenizer("Ingredients : flour , eggs ; milk", " ")))
	fmt.Println(niveaux.NiveauTrois(tokenizer.Tokenizer("Pi is 3.14 and 1,000 is big", " ")))
	fmt.Println(niveaux.NiveauTrois(tokenizer.Tokenizer("e.g. this,that", " ")))
	fmt.Println(niveaux.NiveauTrois(tokenizer.Tokenizer("Punctuation tests are ... kinda boring ,what do you think ?", " ")))
}
