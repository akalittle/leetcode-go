package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zigZagConversion(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		nums int
		want string
	}{
		{
			args: "PAYPALISHIRING",
			nums: 3,
			want: "PAHNAPLSIIGYIR",
		},
		{
			args: "PAYPALISHIRING",
			nums: 4,
			want: "PINALSIGYAHRPI",
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, convert(item.args, item.nums), "输入:%s", item.args, item.nums)
	}
}
