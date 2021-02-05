package trie

const (
	// AlphabetSize total characters in english alphabet
	AlphabetSize = 26
)

<<<<<<< HEAD
type TrieNode struct {
	children  [AlphabetSize]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

// New returns a pointer pointing to a new Trie
func New() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

// Insert inserts a word into the Trie
func (t *Trie) Insert(word string) {
=======
type trieNode struct {
	children  [AlphabetSize]*trieNode
	isWordEnd bool
}

type trie struct {
	root       *trieNode
	nodesCount int
}

// New returns a pointer pointing to a new trie
func New() *trie {
	return &trie{
		root: &trieNode{},
	}
}

// Insert inserts a word into the trie
func (t *trie) Insert(word string) {
>>>>>>> 2896b350131353759aacf42da7047972425068eb
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
<<<<<<< HEAD
			current.children[index] = &TrieNode{}
=======
			current.children[index] = &trieNode{}
			t.nodesCount++
>>>>>>> 2896b350131353759aacf42da7047972425068eb
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

<<<<<<< HEAD
// Contains checks whether the Trie has the given word or not
func (t *Trie) Contains(word string) bool {
=======
// Contains checks whether the trie has the given word or not
func (t *trie) Contains(word string) bool {
>>>>>>> 2896b350131353759aacf42da7047972425068eb
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}

	return current.isWordEnd
<<<<<<< HEAD
}

func (trie *TrieNode, value string, values []string) findWords(){
	for i := 0; i< 26 ; i++ {
		if trie.isWordEnd {
			values.append(value+string(i+'a'))
		}
		if trie.children[i] != nil {
			findWords(trie.children[i], value+string(i+'a'), values)
		}
	}
=======
}

// NodesCount returns number of nodes in the tree.
func (t *trie) NodesCount() int {
	return t.nodesCount
}

// Empty returns true if tree does not contain any nodes except for root
func (t *trie) Empty() bool {
	return t.nodesCount == 0
>>>>>>> 2896b350131353759aacf42da7047972425068eb
}
