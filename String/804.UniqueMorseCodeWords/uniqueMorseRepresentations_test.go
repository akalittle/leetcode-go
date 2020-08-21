package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []string
		want int
	}{
		{
			args: []string{"gin", "zen", "gig", "msg"},
			want: 2,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, uniqueMorseRepresentations(item.args), "输入:%s", item.args)
	}
}
