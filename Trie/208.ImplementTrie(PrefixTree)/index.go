package Trie

//Implement a trie with insert, search, and startsWith methods.
//
//Example:
//
//Trie trie = new Trie()
//
//trie.insert("apple")
//trie.search("apple") // returns true
//trie.search("app")     // returns false
//trie.startsWith("app") // returns true
//trie.insert("app")
//trie.search("app") // returns true
//Note:
//
//You may assume that all inputs are consist of lowercase letters a-z.
//All inputs are guaranteed to be non-empty strings.

type Trie struct {
	value  string
	next   map[string]*Trie
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{value: "", next: make(map[string]*Trie), isWord: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, w := range word {
		n := root.match(string(w))
		if n == nil {
			newNode := &Trie{value: string(w), next: make(map[string]*Trie)}
			root.next[string(w)] = newNode
			root = root.next[string(w)]
		} else {
			root = n
		}
	}
	root.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if n := this.match(word); n != nil && n.isWord {
		return true
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if n := this.match(prefix); n != nil {
		return true
	} else {
		return false
	}
}

func (this *Trie) match(word string) *Trie {
	root := this
	for _, w := range word {
		if _, ok := root.next[string(w)]; ok {
			root = root.next[string(w)]
		} else {
			return nil
		}
	}
	return root
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
