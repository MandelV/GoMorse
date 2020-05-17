package gomorse

import "testing"

func isValid(test *testing.T, tree *Tree, code, letter string) {
	if tree.search(code).Letter != letter {
		test.Error(code, letter, "are not Match")
	}else {
		test.Log(code, letter, "are Matching")
	}
}
func TestEncode(test *testing.T) {

	tree := MorseTree
	isValid(test, tree, ".", "E")
	isValid(test, tree, ".-", "A")
	isValid(test, tree, ".-.", "R")
	isValid(test, tree, ".--", "W")
	isValid(test, tree, ".--.", "P")
	isValid(test, tree, ".---", "J")
	isValid(test, tree, ".-..", "L")
	isValid(test, tree, "..", "I")
	isValid(test, tree, "..-", "U")
	isValid(test, tree, "..-.", "F")
	isValid(test, tree, "...", "S")
	isValid(test, tree, "...-", "V")
	isValid(test, tree, "....", "H")
	isValid(test, tree, "-", "T")
	isValid(test, tree, "-.", "N")
	isValid(test, tree, "-.-", "K")
	isValid(test, tree, "-.--", "Y")
	isValid(test, tree, "-.-.", "C")
	isValid(test, tree, "-..", "D")
	isValid(test, tree, "-..-", "X")
	isValid(test, tree, "-...", "B")
	isValid(test, tree, "--", "M")
	isValid(test, tree, "--.", "G")
	isValid(test, tree, "--.-", "Q")
	isValid(test, tree, "--..", "Z")
	isValid(test, tree, "---", "O")
}
