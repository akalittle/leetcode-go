package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//1->2->3->4->5

func Test_combinationSum(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want: [][]int{
				{7},
				{2, 2, 3},
			},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			want: [][]int{
				{3, 5},
				{2, 3, 3},
				{2, 2, 2, 2},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, combinationSum(item.candidates, item.target), "输入:%s", item.candidates, item.target)
	}
}
