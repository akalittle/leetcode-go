package _47_TopKFrequentElements

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]

func Test_topKFrequent(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		k    int
		want []int
	}{
		{
			args: []int{
				1, 1, 1, 2, 2, 3,
			},
			k:    2,
			want: []int{1, 2},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, topKFrequent(item.args, item.k), "输入:%s", item.args)
	}
}
