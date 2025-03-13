//Node represent each character

type Node struct {

//this is a single letter stored for example letter a,b,c,d,etc

Char string


//store all children of a node

//that is from a-z

//a slice of Nodes(and each child will also have 26 children)

Children [26]*Node

}


/// NewNode this will be used to initialize a new node with 26 children

///each child should first be initialized to nil

func NewNode(char string) *Node {

node := &Node{Char: char}

for i := 0; i < 26; i++ {

node.Children[i] = nil

}

return node

}






// Trie is our actual tree that will hold all of our nodes

//the Root node will be nil

type Trie struct {

RootNode *Node

}


// NewTrie Creates a new trie with a root('constructor')

func NewTrie() *Trie {

//we will not use this node so


//it can be anything

root := NewNode("\000")

return &Trie{RootNode: root}}


 





//Insert inserts a word to the trie

func (t *Trie) Insert(word string) error {

///this will keep track of our current node

///when transversing our tree

///it should always start at the top of our tree

///i.e our root

current := t.RootNode

///remove all spaces from the word

///and convert it to lowercase

strippedWord := 


strings.ToLower(strings.ReplaceAll(word, " ", ""))

for i := 0; i < len(strippedWord); i++ {

//from the ascii table a represent decimal number 97

//from the ascii table b represent decimal number 98

//from the ascii table c represent decimal number 99

/// and so on so basically if we were to say 98-97=1 which means the index of b is 1 and for c is 99-97

///that what is happening below (we are taking the decimal representation of a character and subtracting decimal representation of a)

index := strippedWord[i] - 'a'

///check if current already has a node created at our current node

//if not create the node

if current.Children[index] == nil {

current.Children[index] = 


NewNode(string(strippedWord[i]))

}

current = current.Children[index]

//since we want to support autocomplete

}

return nil

}







// SearchWord will return false if a word we 

//are searching for is not in the trie

//and true otherwise

func (t *Trie) SearchWord(word string) bool {


strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

current := t.RootNode

for i := 0; i < len(strippedWord); i++ {

index := strippedWord[i] - 'a'

//we have encountered null in the path we were transversing meaning this is the last node

///that means this word is not indexed(present) in this trie

if current == nil || current.Children[index] == nil {

return false

}

}

return true

}
