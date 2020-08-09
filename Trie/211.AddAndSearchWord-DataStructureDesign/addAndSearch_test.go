package Trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SymmetricTree(t *testing.T) {
	ast := assert.New(t)

	//tests := []struct {
	//	want bool
	//}{
	//	{
	//		want: false,
	//	},
	//}
	//
	//for _ , := range tests {
	//
	//
	//}

	obj := Constructor()
	obj.AddWord("bad")

	obj.AddWord("dad")
	obj.AddWord("mad")

	ast.Equal(false, obj.Search("pad"))
	ast.Equal(true, obj.Search("bad"))
	ast.Equal(true, obj.Search(".ad"))
	ast.Equal(true, obj.Search("b.."))

}
