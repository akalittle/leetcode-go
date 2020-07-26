package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//matrix = [
//[1,2,3,4],
//[5,1,2,3],
//[9,5,1,2]
//]

func Test_IsToeplitzMatrix(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args [][]int
		want bool
	}{
		{
			args: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			want: true,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, IsToeplitzMatrix(item.args), "输入:%s", item.args)
	}
}
