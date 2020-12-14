package LCOF

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minArray(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{3, 4, 5, 1, 2},
			want: 5,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, minArray(item.args), "输入:%s", item.args)
	}
}
