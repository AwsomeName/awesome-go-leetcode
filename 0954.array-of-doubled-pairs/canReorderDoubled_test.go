package problem0954

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []int
}

type ans struct {
	ans bool
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
				one: []int{3, 1, 3, 6},
			},
			a: ans{
				false,
			},
		},
		question{
			p: para{
				one: []int{2, 1, 2, 6},
			},
			a: ans{
				false,
			},
		},
		question{
			p: para{
				one: []int{4, -2, 2, -4},
			},
			a: ans{
				true,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 4, 16, 8, 4},
			},
			a: ans{
				false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, canReorderDoubled(p.one), "input:%v", p)
	}
}
