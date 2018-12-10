package problem0956

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []int
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
				one: []int{1, 2, 3, 6},
			},
			a: ans{
				6,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 4, 5, 6},
			},
			a: ans{
				10,
			},
		},
		question{
			p: para{
				one: []int{1, 2},
			},
			a: ans{
				0,
			},
		},
		question{
			p: para{
				one: []int{2, 4, 8, 16},
			},
			a: ans{
				0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, tallestBillboard(p.one), "input:%v", p)
	}
}
