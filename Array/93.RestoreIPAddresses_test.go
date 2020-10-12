package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		want []string
	}{
		{
			args: "101023",
			want: []string{
				"1.0.10.23",
				"1.0.102.3",
				"10.1.0.23",
				"10.10.2.3",
				"101.0.2.3",
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, restoreIpAddresses(item.args), "输入:%s", item.args)
	}
}
