package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		want []string
	}{
		{
			args: "23",
			want: []string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, letterCombinations(item.args), "输入:%s", item.args)
	}
}
