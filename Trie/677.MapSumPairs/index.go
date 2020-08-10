package Trie

//Implement a MapSum class with insert, and sum methods.
//
//For the method insert, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value. If the key already existed, then the original key-value pair will be overridden to the new one.
//
//For the method sum, you'll be given a string representing the prefix, and you need to return the sum of all the pairs' value whose key starts with the prefix.
//
//Example 1:
//Input: insert("apple", 3), Output: Null
//Input: sum("ap"), Output: 3
//Input: insert("app", 2), Output: Null
//Input: sum("ap"), Output: 5
type Trie struct {
	key  string
	val  int
	next map[string]*Trie
}

/** Initialize your data structure here. */

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string, v int) {
	root := this
	for _, w := range word {
		n := root.match(string(w))
		if n == nil {
			newNode := &Trie{key: string(w), next: make(map[string]*Trie)}
			root.next[string(w)] = newNode
			root = root.next[string(w)]
		} else {
			root = n
		}
	}

	root.val = v

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

type MapSum struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		trie: &Trie{
			next: make(map[string]*Trie),
		},
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.trie.Insert(key, val)

}

func (this *MapSum) Sum(prefix string) int {
	cur := this.trie
	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return 0
		}
		cur = cur.next[c]
	}

	return sum(cur)
}

func sum(n *Trie) int {
	res := n.val
	for str := range n.next {
		res += sum(n.next[str])
	}
	return res
}
