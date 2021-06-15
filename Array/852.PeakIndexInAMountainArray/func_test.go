package _52_PeakIndexInAMountainArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int

		want int
	}{
		{
			args: []int{
				24, 69, 100, 99, 79, 78, 67, 36, 26, 19,
			},
			want: 2,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, peakIndexInMountainArray(item.args), "输入:%s", item.args)
	}
}
