package gomorse

import (
	"log"
	"testing"
)

func isValid(t *testing.T, tree *Tree, code, letter string) {
	if tree.search(code).Letter != letter {
		t.Error(code, letter, "are not Match")
	}else {
		t.Log(code, letter, "are Matching")
	}
}
//Test each code and its matching letter
func TestInit(t *testing.T) {
	tree := morseTree
	isValid(t, tree, ".", "E")
	isValid(t, tree, ".-", "A")
	isValid(t, tree, ".-.", "R")
	isValid(t, tree, ".--", "W")
	isValid(t, tree, ".--.", "P")
	isValid(t, tree, ".---", "J")
	isValid(t, tree, ".-..", "L")
	isValid(t, tree, "..", "I")
	isValid(t, tree, "..-", "U")
	isValid(t, tree, "..-.", "F")
	isValid(t, tree, "...", "S")
	isValid(t, tree, "...-", "V")
	isValid(t, tree, "....", "H")
	isValid(t, tree, "-", "T")
	isValid(t, tree, "-.", "N")
	isValid(t, tree, "-.-", "K")
	isValid(t, tree, "-.--", "Y")
	isValid(t, tree, "-.-.", "C")
	isValid(t, tree, "-..", "D")
	isValid(t, tree, "-..-", "X")
	isValid(t, tree, "-...", "B")
	isValid(t, tree, "--", "M")
	isValid(t, tree, "--.", "G")
	isValid(t, tree, "--.-", "Q")
	isValid(t, tree, "--..", "Z")
	isValid(t, tree, "---", "O")
}

func TestDecode(t *testing.T) {
	//msg := "DECODEMORSE"
	morse := "-.. . -.-. --- -.. . -- --- .-. ... ."
	plain :=  *Decode(&morse)
	if plain != "DECODEMORSE" {
		t.Error(morse, plain, "Not match")
	}
}

func TestNode_Browse(t *testing.T) {
	if node := morseTree.Groot.Browse("Q"); node == nil {
		t.Error("Node is nil")
	}else if node.Letter != "Q" {
		t.Error("Node not found")
	}
}

func TestTree_GetPath(t *testing.T) {

	log.Println("CODE : ",morseTree.GetPath("Y"))

}