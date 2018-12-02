package problem0951

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []int
}

type ans struct {
	ans string
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
				one: []int{1,2,3,4},
			},
			a: ans{
                "23:41",
			},
		},
		question{
			p: para{
				one: []int{5,5,5,5},
			},
			a: ans{
				"",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, largestTimeFromDigits(p.one), "input:%v", p)
	}
}
