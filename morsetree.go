package gomorse

import "strings"

type Node struct {
	Dot    *Node
	Dash   *Node
	Letter string
}

type Tree struct {
	Groot *Node
}

func (node *Node) insert(code, letter string) {
	if node == nil || len(code) != 1 {
		return
	}
	if code == "." {
		if node.Dot == nil {
			node.Dot = &Node{Letter: letter}
		} else {
			node.Dot.insert(code, letter)
		}
	} else {
		if node.Dash == nil {
			node.Dash = &Node{Letter: letter}
		} else {
			node.Dash.insert(code, letter)
		}
	}
}

func (tree *Tree) search(code string) *Node {
	currentNode := tree.Groot
	for _, partialCode := range code {
		pc := string(partialCode)

		if pc == "." {
			if currentNode.Dot == nil {
				continue
			}
			currentNode = currentNode.Dot

		} else {
			if currentNode.Dash == nil {
				continue
			}
			currentNode = currentNode.Dash

		}
	}
	return currentNode
}

func (tree *Tree) insert(code, letter string) *Tree {
	if tree.Groot == nil {
		tree.Groot = &Node{Letter: letter, Dot: nil, Dash: nil}
	} else {
		if len(code) == 1 {
			tree.Groot.insert(code, letter)
		} else {
			findedNode := tree.search(code)
			findedNode.insert(code[len(code)-1:], letter)
		}

	}

	return tree
}

func initTree() *Tree {
	tree := &Tree{Groot: &Node{Letter: "ROOT"}}

	tree.insert(".", "E").
		insert(".-", "A").
		insert(".-.", "R").
		insert(".--", "W").
		insert(".--.", "P").
		insert(".---", "J").
		insert(".-..", "L").
		insert("..", "I").
		insert("..-", "U").
		insert("..-.", "F").
		insert("...", "S").
		insert("...-", "V").
		insert("....", "H").
		insert("-", "T").
		insert("-.", "N").
		insert("-.-", "K").
		insert("-.--", "Y").
		insert("-.-.", "C").
		insert("-..", "D").
		insert("-..-", "X").
		insert("-...", "B").
		insert("--", "M").
		insert("--.", "G").
		insert("--.-", "Q").
		insert("--..", "Z").
		insert("---", "O")


	return tree
}

type Callback func(string) bool
func (node *Node) Browse(letter string) *Node{
	/*if node != nil {
		n := node.Dot.Browse(letter)
		if node.Letter == letter {
			return n
		}
		node.Dash.Browse(letter)
	}*/
	return nil
}

//Encode message to morse
func  Encode(message *string) (morse *string){




	return nil
}

//Decode morse to message
func Decode(morse *string) (message *string){
	splited := strings.Split(*morse, " ")
	msg := ""

	for _, morseLetter := range splited {
		msg += morseTree.search(morseLetter).Letter
	}
	return &msg
}
//MorseTree Binary tree that represent the morse language
var morseTree *Tree = initTree()
