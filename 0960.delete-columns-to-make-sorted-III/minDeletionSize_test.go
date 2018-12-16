package problem0955

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []string
}

type ans struct {
	ans int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []string{"babca", "bbazb"},
			},
			a: ans{
				3,
			},
		},
		question{
			p: para{
				one: []string{"edcba"},
			},
			a: ans{
				4,
			},
		},
		question{
			p: para{
				one: []string{"ghi", "def", "abc"},
			},
			a: ans{
				0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, minDeletionSize(p.one), "input:%v", p)
	}
}
