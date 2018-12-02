package problem0950

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []int
}

type ans struct {
	ans []int
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
				one: []int{17,13,11,2,3,5,7},
			},
			a: ans{
                []int{2,13,3,11,5,17,7},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, deckRevealedIncreasing(p.one), "input:%v", p)
	}
}
