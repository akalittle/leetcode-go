package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		s1   string
		s2   string
		want []int
	}{
		{
			s1:   "cbaebabacd",
			s2:   "abc",
			want: []int{0, 6},
		},
		{
			s1:   "baa",
			s2:   "aa",
			want: []int{1},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, findAnagrams(item.s1, item.s2), "输入:%s", item.s1)
	}
}
