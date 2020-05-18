package gomorse

import (
	"strings"
)

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


func (node *Node) path(path *[]string, letter string) bool {
	if node == nil {
		return false
	}
	if node.Dot.path(path, letter) {
		*path = append(*path, ".")
		return true
	}else if node.Dash.path(path, letter){
		*path = append(*path, "-")
		return true
	} else if node.Letter == letter{
		return true
	}else{
		return false
	}
}

func (tree *Tree)GetPath(letter string) (code string){
	var path []string
	morseTree.Groot.path(&path, letter)

	return func() string{
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1{
			path[i], path[j] = path[j], path[i]
		}
		code = strings.Join(path, "")
		return code
	}()
}
// Browse enable to find the node that contains the given letter
func (node *Node) Browse(letter string) (NodeFound *Node){
	if node == nil {
		return nil
	}
	 nDot, nDash := node.Dot.Browse(letter), node.Dash.Browse(letter)

	 if nDot != nil && nDot.Letter == letter{
		 return nDot
	 }else if nDash != nil && nDash.Letter == letter{

		 return nDash
	 }
	return node
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
