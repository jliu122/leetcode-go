package trie

// 208
type Trie struct {
	root TrieNode
}

type TrieNode struct {
	IsEnd bool
	Val rune
	Children []*TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{TrieNode{false, -1, make([]*TrieNode, 26)}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {

	curr := this.root
	for _, c := range word {
		index := c - 'a'
		if curr.Children[index] == nil {
			curr.Children[index] = &TrieNode{false, c, make([]*TrieNode, 26)}
		}
		curr = *curr.Children[index]
	}
	curr.IsEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	curr := this.root
	for _, c := range word {
		index := c - 'a'
		if curr.Children[index] == nil {
			return false
		}
		curr = *curr.Children[index]
	}
	return curr.IsEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	curr := this.root
	for _, c := range prefix {
		index := c - 'a'
		if curr.Children[index] == nil {
			return false
		}
		curr = *curr.Children[index]
	}
	return true
}
