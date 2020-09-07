package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		num1 []int
		m    int
		num2 []int
		n    int

		want []int
	}{
		{
			num1: []int{
				1, 2, 3, 0, 0, 0,
			},
			m: 3,
			num2: []int{
				2, 5, 6,
			},
			n: 3,
			want: []int{
				1, 2, 2, 3, 5, 6,
			},
		},
	}

	for _, item := range tests {
		merge(item.num1, item.m, item.num2, item.n)
		ast.Equal(item.want, item.num1, "输入:%s", item.num1)
	}
}
