package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Given a string, find the first non-repeating character in it and return its index.If it doesn't exist, return -1.
//
//Examples:
//
//s = "leetcode"
//return 0.
//
//s = "loveleetcode"
//return 2.
// 
//
//Note: You may assume the string contains only lowercase English letters.

func Test_firstUniqChar(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		want int
	}{
		{
			args: "leetcode",
			want: 0,
		},
		{
			args: "loveleetcode",
			want: 2,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, firstUniqChar(item.args), "输入:%s", item.args)
	}
}
