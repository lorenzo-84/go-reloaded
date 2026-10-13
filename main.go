package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . fichier_entree fichier_sortie")
		return
	}

	fichierEntree := os.Args[1]
	fichierSortie := os.Args[2]

	data, err := os.ReadFile(fichierEntree)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	texte := string(data)
	resultat := Process(texte)

	err = os.WriteFile(fichierSortie, []byte(resultat), 0644)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
}
