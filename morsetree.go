package gomorse

import (
	"errors"
	"strings"
	"sync"
)
//MorseTree Binary tree that represent the morse language
var morseTree *Tree = initTree()
type Word struct {
	index int
	word string
}
type Node struct {
	Dot    *Node
	Dash   *Node
	Letter string
}

type Tree struct {
	Groot *Node
}
//Init the morse tree
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
//insert a new node at the parent node depending on whether code is dot (.) or dash (-)
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

//search a node with given morse code
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

//insert a new node at the end of the path describe by the morse code
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

//Get the path of the given letter et put dot or slash on the path string array
func (node *Node) path(path *[]string, letter string) bool {
	if path == nil {
		panic("path is nil")
	}
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

//browse enable to find the node that contains the given letter
func (node *Node) browse(letter string) (NodeFound *Node){
	if node == nil {
		return nil
	}
	 nDot, nDash := node.Dot.browse(letter), node.Dash.browse(letter)

	 if nDot != nil && nDot.Letter == letter{
		 return nDot
	 }else if nDash != nil && nDash.Letter == letter{

		 return nDash
	 }
	return node
}

//Encode message to morse
func Encode(message *string) (morse *string, err error){

	*message = strings.ToUpper(*message)
	code := ""
	morse = &code

	for _, letter := range *message{
		l := string(letter)
		if l == " " {
			code += "/"
			continue
		}
		if mCode, er := morseTree.GetCode(l); er == nil {
			code += mCode
			code += " "
		}else{
			return nil, er
		}
	}
	//Just beautify the code
	code = strings.Replace(code, " /", "/", len(code))
	if code[len(code)-1:] == " "{
		code = code[:len(code)-1]
	}
	return morse, nil
}

//decodeWord will be use as go routine to translate morse word to plainWord
func decodeWord(wg *sync.WaitGroup, index int , out *Word,  word string){
	defer wg.Done()
	w := ""
	for _, code := range strings.Split(word, " ") {
		letter, _ := morseTree.GetLetter(code)
		w += letter
	}
	*out = Word{index: index, word: w}
}
//Decode morse to message
func Decode(morse *string) (message *string, err error) {
	if morse == nil {
		return nil, errors.New("morse is nil")
	}
	msg := ""
	message = &msg

	wordsStr := strings.Split(*morse, "/")

	out := make([]Word, len(wordsStr))
	wg := new(sync.WaitGroup)

	//prepare go routine to work on each word of the morse code
	for i, word := range wordsStr{
		wg.Add(1)
		go decodeWord(wg, i, &out[i], word)
	}
	wg.Wait()
	for _, v := range out {
		msg += v.word + " "
	}
	if (*morse)[len(*morse)-1:] == " "{
		*morse = (*morse)[:len(*morse)-1]
	}
	if msg[len(msg)-1:] == " "{
		msg = msg[:len(msg)-1]
	}
	return message, nil
}

//GetLetter morse to message
func (tree *Tree) GetLetter(morse string) (letter string, err error){
	if node := tree.search(morse); node != nil {
		return node.Letter, nil
	}
	return "", errors.New("not found")
}

// GetCode get the code of the given letter
func (tree *Tree) GetCode(letter string) (code string, err error){
	var path []string
	morseTree.Groot.path(&path, letter)

	//Reverse the path
	return func() (string, error){
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1{
			path[i], path[j] = path[j], path[i]
		}
		code = strings.Join(path, "")
		if code == "" {
			return "", errors.New("code not found for : " + "\"" + letter + "\"")
		}else {
			return code, nil
		}
	}()
}