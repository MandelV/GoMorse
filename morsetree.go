package gomorse

import (
	"errors"
	"strings"
	"sync"
)

//Node represent a letter in Tree
type Node struct {
	Dot    *Node
	Dash   *Node
	Letter string
}

//Tree represents the morse language
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

	//ADD NUMBERS
	tree.insert(".----", "1")
	tree.insert("..---", "2")
	tree.insert("...--", "3")
	tree.insert("....-", "4")
	tree.insert(".....", "5")
	tree.insert("-....", "6")
	tree.insert("--...", "7")
	tree.insert("---..", "8")
	tree.insert("----.", "9")
	tree.insert("-----", "0")




	return tree
}

//insert a new node at the parent node depending on whether code is dot (.) or dash (-)
func (node *Node) insert(code, letter string) {
	if node == nil || len(code) != 1 {
		return
	}
	if code == "." { //DOT
		if node.Dot == nil {
			node.Dot = &Node{Letter: letter}
		}else {
			node.Dot.insert(code, letter)
		}
	} else {// DASH
		if node.Dash == nil {
			node.Dash = &Node{Letter: letter}
		} else {
			node.Dash.insert(code, letter)
		}
	}
}

// hasNext return double bool dot and dash that are true if child node is presents
func (node *Node) hasNext() (dot bool, dash bool){
	return node.Dot != nil, node.Dash != nil
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

//addPathNodes Creates each node from the racine to the last node following the code
func (tree *Tree) addPathNodes(code string) *Node{

	currentNode := tree.Groot
	for _, partialCode := range code {
		pc := string(partialCode)

		if pc == "." {
			if currentNode.Dot == nil {
				currentNode.Dot = &Node{Letter: "", Dash: nil, Dot: nil}
			}
			currentNode = currentNode.Dot

		} else {
			if currentNode.Dash == nil {
				currentNode.Dash = &Node{Letter: "", Dash: nil, Dot: nil}
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
			foundNode := tree.addPathNodes(code)
			//if the node has no child and no letter setted then we estimate that is the right node to put the letter in it
			if dot, dash := foundNode.hasNext(); !dot && !dash && foundNode.Letter == ""{
				foundNode.Letter = letter
				return tree
			}
			foundNode.insert(code[len(code)-1:], letter)
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

//decodeWord will be used as go routine to translate morse word to plainWord
func decodeWord(wg *sync.WaitGroup, out *string,  word string){
	defer wg.Done()
	w := ""
	for _, code := range strings.Split(word, " ") {
		letter, _ := GetLetter(code)
		w += letter
	}
	*out = w
}



/* ========================================
 *
 *       EXPORTED ELEMENTS
 *
 * ========================================
 */
//MorseTree Binary tree that represent the morse language
var MorseTree = initTree()
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
		if mCode, er := GetCode(l); er == nil {
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

//Decode morse to message
func Decode(morse *string) (message *string, err error) {
	if morse == nil {
		return nil, errors.New("morse is nil")
	}
	msg := ""
	message = &msg

	wordsStr := strings.Split(*morse, "/")

	out := make([]string, len(wordsStr))
	wg := new(sync.WaitGroup)

	//prepare go routine to work on each word of the morse code
	for i, word := range wordsStr{
		wg.Add(1)
		go decodeWord(wg, &out[i], word)
	}
	//Waiting each go routine
	wg.Wait()
	//Create the final message
	for _, v := range out {
		msg += v + " "
	}
	//Beautify the final message
	if (*morse)[len(*morse)-1:] == " "{
		*morse = (*morse)[:len(*morse)-1]
	}
	if msg[len(msg)-1:] == " "{
		msg = msg[:len(msg)-1]
	}
	return message, nil
}

//GetLetter morse to message
func GetLetter(morse string) (letter string, err error){
	if node := MorseTree.search(morse); node != nil {
		return node.Letter, nil
	}
	return "", errors.New("not found")
}

// GetCode get the code of the given letter
func GetCode(letter string) (code string, err error){
	var path []string
	MorseTree.Groot.path(&path, letter)

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
