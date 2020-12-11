package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		//{
		//	s1:   "ab",
		//	s2:   "eidbaooo",
		//	want: true,
		//},
		{
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, checkInclusion(item.s1, item.s2), "输入:%s", item.s1)
	}
}
