package main

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		entree string
		sortie string
	}{
		{"1E (hex)", "30"},
		{"10 (bin)", "2"},
		{"hello (up)", "HELLO"},
		{"HELLO (low)", "hello"},
		{"hello (cap)", "Hello"},
		{"hello world (up,2)", "HELLO WORLD"},
		{"HELLO WORLD (low,2)", "hello world"},
		{"hello world (cap,2)", "Hello World"},
		{"Bonjour , comment allez vous ?", "Bonjour, comment allez vous?"},
		{"a apple", "an apple"},
	}

	for _, test := range tests {
		resultat := Process(test.entree)

		if resultat != test.sortie {
			t.Errorf("attendu %q, obtenu %q", test.sortie, resultat)
		}
	}
}
