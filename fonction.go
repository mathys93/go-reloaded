package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Vérifier qu'on a bien 2 arguments : fichier d'entrée et fichier de sortie
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . <entree.txt> <sortie.txt>")
		return
	}

	// Lire le fichier d'entrée
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	// Découper le texte en mots
	words := strings.Fields(string(data))

	// Appliquer les règles
	words = convertHex(words)

	// Recoller les mots et écrire le résultat
	result := strings.Join(words, " ")
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}
}

// convertHex remplace le mot avant "(hex)" par sa valeur décimale
// et supprime le marqueur "(hex)"
func convertHex(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			if i > 0 {
				n, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(n, 10)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}