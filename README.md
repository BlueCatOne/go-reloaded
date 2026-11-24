# go-reloaded

Outil en Go qui lit un texte brut passé en premier argument, le corrige selon un ensemble de règles (conversion de nombres, casse, ponctuation, guillemets, article *a/an*) et écrit le résultat dans le fichier passé en deuxième argument, en le créant s'il n'existe pas.

Aucune dépendance externe : uniquement la bibliothèque standard.

## Prérequis

- Go 1.27.1 (voir `go.mod`)

## Utilisation

```bash
go run . <entrée> <sortie>
```

Le programme lit `<entrée>` à la racine du projet et écrit `<sortie>` (créé ou écrasé). Si `<entrée>` est introuvable, ou si les arguments sont invalides, le programme s'arrête avec le code de sortie 1.


## Règles appliquées

Le texte est traité ligne par ligne, mot par mot.

| Niveau | Règle | Exemple |
|:-:|---|---|
| 1 | `(hex)` et `(bin)` convertissent le mot précédent (hexadécimal en majuscules / binaire) en décimal. Une valeur invalide reste inchangée. | `1E (hex)` → `30` · `10 (bin)` → `2` |
| 2 | `(up)`, `(low)`, `(cap)` changent la casse du mot précédent. `(up, n)`, `(low, n)`, `(cap, n)` l'appliquent aux `n` mots précédents (avec l'espace après la virgule). Les balises sont retirées. | `so exciting (up, 2)` → `SO EXCITING` |
| 3 | La ponctuation `. , ! ? : ;` est collée au mot précédent et suivie d'une espace. Les groupes (`...`, `!?`) sont conservés, de même que `3.14` ou `1,000`. | `there ,and` → `there, and` |
| 4 | Les apostrophes simples `'` servant de guillemets sont collées au texte qu'elles entourent. | `' awesome '` → `'awesome'` |
| 5 | `a` devient `an` devant une voyelle minuscule, et inversement (majuscule initiale respectée). | `a amazing` → `an amazing` |

## Exemple

Entrée (`<entrée>`) :

```text
if (cap) you can read this ,your program works . i (up) hid this capsule in 7E0 (hex) , after 101010 (bin) cups of coffee and a unreasonable amount of DEBUGGING (low) .
```

Sortie (`<sortie>`) :

```text
If you can read this, your program works. I hid this capsule in 2016, after 42 cups of coffee and an unreasonable amount of debugging.
```

## Structure

```text
go-reloaded/
├── main.go                  # point d'entrée : lecture, application des niveaux, écriture
├── <entrée>              # texte d'entrée
├── <sortie>               # sortie générée
└── src/
    ├── tokenizer/           # découpe une chaîne selon un séparateur
    ├── niveaux/             # une règle par fichier : niveauUn.go … niveauCinq.go
    ├── manipulationFichier/ # Lire, Ecrire, Verifier
```

## Limites connues

- La casse ne gère que l'ASCII (pas les lettres accentuées), et la règle *a/an* ne tient compte que des voyelles `a e i o u` (pas du *h*).
