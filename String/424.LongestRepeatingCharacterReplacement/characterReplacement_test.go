package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		s1   string
		s2   int
		want int
	}{
		{
			s1:   "ABAB",
			s2:   2,
			want: 4,
		},
		{
			s1:   "AABABBA",
			s2:   1,
			want: 4,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, characterReplacement(item.s1, item.s2), "输入:%s", item.s1)
	}
}
