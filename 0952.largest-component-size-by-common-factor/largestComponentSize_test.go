package problem0952

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
				one: []int{4,6,15,35},
			},
			a: ans{
				4,
			},
		},
		question{
			p: para{
				one: []int{20,50,9,63},
			},
			a: ans{
				2,
			},
		},
		question{
			p: para{
				one: []int{2,3,6,7,4,12,21,39},
			},
			a: ans{
				8,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, largestComponentSize(p.one), "input:%v", p)
	}
}
