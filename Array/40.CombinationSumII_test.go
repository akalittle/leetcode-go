package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//1->2->3->4->5

func Test_combinationSum2(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		//{
		//	candidates: []int{2, 3, 5},
		//	target:     8,
		//	want: [][]int{
		//		{3, 5},
		//		{2, 3, 3},
		//		{2, 2, 2, 2},
		//	},
		//},
	}

	for _, item := range tests {
		ast.Equal(item.want, combinationSum2(item.candidates, item.target), "输入:%s", item.candidates, item.target)
	}
}
