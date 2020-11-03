package LCOF

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findRepeatNumber(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		m    int
		n    int
		k    int
		want int
	}{
		{
			m:    2,
			n:    3,
			k:    1,
			want: 3,
		},
		{
			m:    3,
			n:    1,
			k:    0,
			want: 1,
		},

		{
			m:    23,
			n:    35,
			k:    4,
			want: 15,
		},
		{
			m:    1,
			n:    3,
			k:    11,
			want: 3,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, movingCount(item.m, item.n, item.k), "输入:%s", item.m, item.n, item.k)
	}
}
