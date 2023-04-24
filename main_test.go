package main

import (
	"io/ioutil"
	"testing"
)

func TestMain(t *testing.T) {
	result, _ := ioutil.ReadFile("standard.txt")
	var texte = string(result)

	//teste FoundFile
	if foundFile("standard") != texte {
		t.Errorf("Error")
	}
	//teste CaseResolved
	if !CaseResolved("\\n\\n") {
		t.Errorf("Error")
	}
	//teste ajout ascii
	if ajoutASCII() == nil {
		t.Errorf("Error")
	}
	//teste Position
	var position = []int{0, 2, 1, 4, 3}
	var textePhrase = []string{"Hello", "Name", "My", "Matar", "Is"}
	if foundposition(position, textePhrase) == nil {
		t.Errorf("Error")
	}
	// teste FoundLine
	if foundLine(texte) == nil {
		t.Error("Error")
	}
	//teste comparaison
	var tabascii []rune
	var positions []int
	var tab []int
	for i := ' '; i <= '~'; i++ {
		tabascii = append(tabascii, i)
	}
	for i := 0; i <= 95; i++ {
		tab = append(tab, i)
	}
	if compraison(tabascii, "Hello Guys", positions, tab) == nil {
		t.Errorf("Error")
	}

	//teste SplitTexte
	//Teste horizontale
	textePhrase = []string{" _  | | | | | | |_| (_) 		", "                                                ", "   _  _    	_| || |_     |_  __  _| 	_| || |_     |_  __  _| 	 |_||_|   			  			     "}
	var x [9][]string = SplitTexte(textePhrase)
	if len(x) == 0 {
		t.Errorf("Error")
	}

}
