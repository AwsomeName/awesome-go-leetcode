package problem0946

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []int
	two []int
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
				one: []int{1, 2, 3, 4, 5},
				two: []int{4, 5, 3, 2, 1},
			},
			a: ans{
				true,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 4, 5},
				two: []int{4, 3, 5, 1, 2},
			},
			a: ans{
				false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, validateStackSequences(p.one, p.two), "input:%v", p)
	}
}
