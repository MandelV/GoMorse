package gomorse

import (
	"strings"
	"testing"
)

func isValid(t *testing.T, tree *Tree, code, letter string) {
	if l := tree.search(code).Letter; l != letter {
		t.Error(code, letter, l, "are not Match")
	}
}
//Test each code and its matching letter
func TestInit(t *testing.T) {
	tree := morseTree
	//LETTERS
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

	//NUMBERS
	isValid(t, tree,".----", "1")
	isValid(t, tree,"..---", "2")
	isValid(t, tree,"...--", "3")
	isValid(t, tree,"....-", "4")
	isValid(t, tree,".....", "5")
	isValid(t, tree,"-....", "6")
	isValid(t, tree,"--...", "7")
	isValid(t, tree,"---..", "8")
	isValid(t, tree,"----.", "9")
	isValid(t, tree,"-----", "0")

}

func TestGetLetter(t *testing.T) {
	//msg := "D"
	morse := "-.."
	if plain, err :=  morseTree.GetLetter(morse); err != nil {
		t.Error(morse, plain, "Not match")
	}
}

func TestNode_Browse(t *testing.T) {
	if node := morseTree.Groot.browse("Q"); node == nil {
		t.Error("Node is nil")
	}else if node.Letter != "Q" {
		t.Error("Node not found")
	}
}

func TestTree_GetCode(t *testing.T) {
	if code, err := morseTree.GetCode("Y"); err != nil {
		t.Error("not found : ", code, " != ", "Y", err)
	}
}

func TestEncode(t *testing.T) {
	morse := ".- -... -.-. -.. ./..-. --. .... .. .--- -.-/.-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --../----- .---- ..--- ...-- ....- ..... -.... --... ---.. ----."
	msg := "abcde fghijk lmnopqrstuvwxyz 0123456789"
	if code, err := Encode(&msg); err != nil {
		t.Error(err)
	}else if *code != morse {
		t.Error("\""+*code+"\"","!=", "\""+morse+"\"", "code does not match")
	}
}

func TestDecode(t *testing.T) {
	morse := ".- -... -.-. -.. ./..-. --. .... .. .--- -.-/.-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --../----- .---- ..--- ...-- ....- ..... -.... --... ---.. ----."
	msg := "abcde fghijk lmnopqrstuvwxyz 0123456789"

	if message, err := Decode(&morse); err != nil {
		t.Error(err)
	}else if *message != strings.ToUpper(msg) {
		t.Error("\"" +*message + "\"", "!=", "\""+strings.ToUpper(msg)+"\"", "msg does not match")
	}
}