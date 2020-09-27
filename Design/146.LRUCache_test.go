package Design

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRU(t *testing.T) {
	ast := assert.New(t)

	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	ast.Equal(lru.Get(1), 1)

	lru.Put(3, 3)

	ast.Equal(lru.Get(2), -1)

	lru.Put(4, 4)

	ast.Equal(lru.Get(1), -1)

}
