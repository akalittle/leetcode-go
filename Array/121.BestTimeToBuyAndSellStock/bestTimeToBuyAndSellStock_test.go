package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{
				7, 1, 5, 3, 6, 4,
			},

			want: 5,
		},
		{
			args: []int{
				7, 6, 4, 3, 1,
			},

			want: 0,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, maxProfit(item.args), "输入:%s", item.args)
	}
}
