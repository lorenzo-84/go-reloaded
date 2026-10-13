package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Process(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i := range lignes {
		lignes[i] = traiterLigne(lignes[i])
	}

	return strings.Join(lignes, "\n")
}

func traiterLigne(texte string) string {
	re := regexp.MustCompile(`\((hex|bin|up|low|cap)(,\s*\d+)?\)`)
	texte = re.ReplaceAllString(texte, " $0 ")
	mots := strings.Fields(texte)

	for i, mot := range mots {
		if mot == "(hex)" || mot == "(bin)" {
			base := 16
			if mot == "(bin)" {
				base = 2
			}

			if n, err := strconv.ParseInt(mots[i-1], base, 64); err == nil {
				mots[i-1] = strconv.FormatInt(n, 10)
			}

			mots[i] = ""
			continue
		}

		if strings.HasPrefix(mot, "(up") ||
			strings.HasPrefix(mot, "(low") ||
			strings.HasPrefix(mot, "(cap") {

			n := getNumber(mot)

			for j := 1; j <= n && i-j >= 0; j++ {
				switch {
				case strings.HasPrefix(mot, "(up"):
					mots[i-j] = strings.ToUpper(mots[i-j])
				case strings.HasPrefix(mot, "(low"):
					mots[i-j] = strings.ToLower(mots[i-j])
				default:
					mots[i-j] = capitalize(mots[i-j])
				}
			}

			mots[i] = ""
		}
	}

	resultat := []string{}
	for _, mot := range mots {
		if mot != "" {
			resultat = append(resultat, mot)
		}
	}

	texte = strings.Join(resultat, " ")
	texte = traiterPonctuation(texte)
	texte = traiterApostrophes(texte)

	return traiterArticles(texte)
}

func capitalize(mot string) string {
	if mot == "" {
		return mot
	}

	r := []rune(strings.ToLower(mot))
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func getNumber(mot string) int {
	debut := strings.Index(mot, ",")
	fin := strings.Index(mot, ")")

	if debut == -1 || fin == -1 {
		return 1
	}

	n, err := strconv.Atoi(strings.TrimSpace(mot[debut+1 : fin]))
	if err != nil {
		return 1
	}

	return n
}

func traiterPonctuation(texte string) string {
	re := regexp.MustCompile(`\s*([.,!?;:]+)\s*`)
	return strings.TrimSpace(re.ReplaceAllString(texte, "$1 "))
}

func traiterApostrophes(texte string) string {
	re := regexp.MustCompile(`\s*'\s*`)
	return re.ReplaceAllString(texte, "'")
}

func traiterArticles(texte string) string {
	mots := strings.Fields(texte)

	for i := 0; i < len(mots)-1; i++ {
		if strings.ToLower(mots[i]) == "a" &&
			strings.ContainsRune("aeiouh", unicode.ToLower(rune(mots[i+1][0]))) {
			mots[i] = "an"
		}
	}

	return strings.Join(mots, " ")
}
