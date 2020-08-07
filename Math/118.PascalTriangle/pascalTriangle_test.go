package Math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PascalTriangle(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want [][]int
	}{
		{
			args: 1,
			want: [][]int{{1}},
		},
		{
			args: 3,
			want: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},

		//	[1],
		//	[1,1],
		//[1,2,1],
		//[1,3,3,1],
		//[1,4,6,4,1]
		{
			args: 5,
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, generate(item.args), "输入:%s", item.args)
	}
}
