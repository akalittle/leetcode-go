package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{
				1, 1, 1, 2, 2, 3,
			},
			want: 3,
		},
		{
			args: []int{
				1, 1,
			},
			want: 1,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, removeDuplicates(item.args), "输入:%s", item.args)
	}
}
