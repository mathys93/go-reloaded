package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	words = convertBin(words)
	words = convertUp(words)
	words = convertLow(words)
	words = convertCap(words)

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

// convertBin remplace le mot avant "(bin)" par sa valeur décimale
// et supprime le marqueur "(bin)"
func convertBin(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" {
			if i > 0 {
				n, err := strconv.ParseInt(words[i-1], 2, 64)
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

// convertUp met en majuscules le mot avant "(up)",
// ou les n mots avant "(up, n)", puis supprime le marqueur
func convertUp(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(up," && i+1 < len(words) {
			n, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err == nil {
				for j := i - 1; j >= 0 && j >= i-n; j-- {
					words[j] = strings.ToUpper(words[j])
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
	}
	return words
}

// convertLow met en minuscules le mot avant "(low)",
// ou les n mots avant "(low, n)", puis supprime le marqueur
func convertLow(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(low," && i+1 < len(words) {
			n, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err == nil {
				for j := i - 1; j >= 0 && j >= i-n; j-- {
					words[j] = strings.ToLower(words[j])
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
	}
	return words
}

// capitalize met la première lettre en majuscule et le reste en minuscules
// exemple : "bRIDGE" -> "Bridge"
func capitalize(s string) string {
	r := []rune(strings.ToLower(s))
	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}
	return string(r)
}

// convertCap met une majuscule au mot avant "(cap)",
// ou aux n mots avant "(cap, n)", puis supprime le marqueur
func convertCap(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" {
			// Cas simple : un seul mot
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(cap," && i+1 < len(words) {
			// Cas avec nombre : "(cap," puis "2)" dans le mot suivant
			n, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err == nil {
				for j := i - 1; j >= 0 && j >= i-n; j-- {
					words[j] = capitalize(words[j])
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
	}
	return words
}