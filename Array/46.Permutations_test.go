package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		candidates []int
		want       [][]int
	}{
		{
			candidates: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
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
		ast.Equal(item.want, permute(item.candidates), "输入:%s", item.candidates)
	}
}
