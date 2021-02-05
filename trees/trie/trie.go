package trie

const (
	// AlphabetSize total characters in english alphabet
	AlphabetSize = 26
)

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
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

// Contains checks whether the Trie has the given word or not
func (t *Trie) Contains(word string) bool {
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

func (trie *TrieNode, value string, values []string) findWords(){
	for i := 0; i< 26 ; i++ {
		if trie.isWordEnd {
			values.append(value+string(i+'a'))
		}
		if trie.children[i] != nil {
			findWords(trie.children[i], value+string(i+'a'), values)
		}
	}
}
