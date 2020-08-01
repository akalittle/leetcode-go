package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intersection(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			arr1: []int{
				1, 2, 2, 1,
			},
			arr2: []int{
				2, 2,
			},
			want: []int{
				2,
			},
		},
		{
			arr1: []int{
				4, 9, 5,
			},
			arr2: []int{
				9, 4, 9, 8, 4,
			},
			want: []int{
				9, 4,
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, intersection(item.arr1, item.arr2), "输入:%s", item.arr1, item.arr2)
	}
}
