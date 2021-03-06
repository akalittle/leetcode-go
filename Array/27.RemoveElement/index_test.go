package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		val  int
		want int
	}{
		{
			args: []int{
				1, 1, 1, 2, 2, 3,
			},
			val:  3,
			want: 5,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, removeElement(item.args, item.val), "输入:%s", item.args)
	}
}
