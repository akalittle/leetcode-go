package Trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mapSumPairs(t *testing.T) {

	ast := assert.New(t)

	obj := Constructor()

	obj.Insert("apple", 3)

	ast.Equal(3, obj.Sum("ap"))

	obj.Insert("app", 2)

	ast.Equal(5, obj.Sum("ap"))



}
