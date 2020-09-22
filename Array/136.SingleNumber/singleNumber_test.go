package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int

		want int
	}{
		{
			args: []int{
				2, 2, 1,
			},
			want: 1,
		},
		{
			args: []int{
				4, 1, 2, 1, 2,
			},
			want: 4,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, singleNumber(item.args), "输入:%s", item.args)
	}
}
