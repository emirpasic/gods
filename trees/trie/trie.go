package trie

const (
	// AlphabetSize total characters in english alphabet
	AlphabetSize = 26
)

type trieNode struct {
	children  [AlphabetSize]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

// New returns a pointer pointing to a new trie
func New() *trie {
	return &trie{
		root: &trieNode{},
	}
}

// Insert inserts a word into the trie
func (t *trie) Insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

// Contains checks whether the trie has the given word or not
func (t *trie) Contains(word string) bool {
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
}
