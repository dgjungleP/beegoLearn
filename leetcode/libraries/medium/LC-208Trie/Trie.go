package main

import "fmt"

type Trie struct {
	endNode bool
	next    map[string]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{endNode: false,
		next: make(map[string]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, chartV := range word {
		next, ok := root.next[string(chartV)]
		if !ok {
			newRoot := Constructor()
			root.next[string(chartV)] = &newRoot
			root = &newRoot
		} else {
			root = next
		}
	}
	root.endNode = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for _, chartV := range word {
		next, ok := root.next[string(chartV)]
		if !ok {
			return false
		}
		root = next
	}
	return root.endNode
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for _, chartV := range prefix {
		next, ok := root.next[string(chartV)]
		if !ok {
			return false
		}
		root = next
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	word := "apple"
	prefix := "app"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search(word)
	param_3 := obj.StartsWith(prefix)
	fmt.Println(param_2, param_3)
}
