package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []string
		want string
	}{
		{
			args: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			args: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			args: []string{},
			want: "",
		},
		{
			args: []string{"", "a", "ab"},
			want: "",
		},
		{
			args: []string{"a", "", "ab"},
			want: "",
		},
		{
			args: []string{"c", "c"},
			want: "c",
		},
		{
			args: []string{"ca", "a"},
			want: "",
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, longestCommonPrefix(item.args), "输入:%s", item.args)
	}
}
