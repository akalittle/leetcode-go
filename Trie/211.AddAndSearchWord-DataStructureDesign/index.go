package Trie

//Design a data structure that supports the following two operations:
//
//void addWord(word)
//bool search(word)
//search(word) can search a literal word or a regular expression string containing only letters a-z or..A.means it can represent any one letter.
//
//Example:
//
//addWord("bad")
//addWord("dad")
//addWord("mad")
//search("pad") -> false
//search("bad") -> true
//search(".ad") -> true
//search("b..") -> true
//Note:
//You may assume that all words are consist of lowercase letters a-z.

type Trie struct {
	value  string
	next   map[string]*Trie
	isWord bool
}

/** Initialize your data structure here. */

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

type WordDictionary struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		trie: &Trie{
			next: make(map[string]*Trie),
		},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.match(this.trie, word, 0)
}

func (this *WordDictionary) match(t *Trie, word string, index int) bool {
	if index == len(word) {
		return t.isWord
	}

	c := string([]rune(word)[index])

	if c != "." {
		if t.next[c] == nil {
			return false
		}
		return this.match(t.next[c], word, index+1)
	} else {
		for w := range t.next {
			if this.match(t.next[w], word, index+1) {
				return true
			}
		}
		return false
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
