package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addStrings(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{
			num1: "111",
			num2: "222",
			want: "333",
		},
		{
			num1: "911",
			num2: "222",
			want: "1133",
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, addStrings(item.num1, item.num2), "输入:%s", item.num1, item.num2)
	}
}
