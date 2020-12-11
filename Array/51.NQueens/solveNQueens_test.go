package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want [][]string
	}{
		{
			args: 4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, solveNQueens(item.args), "输入:%s", item.args)
	}
}
