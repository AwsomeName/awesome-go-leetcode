package problem0900

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one string
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
				one: "awsome",
			},
			a: ans{
                23,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, testFunc(p.one), "input:%v", p)
	}
}
