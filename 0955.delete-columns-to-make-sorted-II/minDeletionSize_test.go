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
				one: []string{"ca", "bb", "ac"},
			},
			a: ans{
				1,
			},
		},
		question{
			p: para{
				one: []string{"xc", "yb", "za"},
			},
			a: ans{
				0,
			},
		},
		question{
			p: para{
				one: []string{"ax", "ay", "ba","bb"},
			},
			a: ans{
				0,
			},
		},
		question{
			p: para{
				one: []string{"zyx", "wvu", "tsr"},
			},
			a: ans{
				3,
			},
		},
		question{
			p: para{
				one: []string{"xga", "xfb", "yfa"},
			},
			a: ans{
				1,
			},
		},
		question{
			p: para{
				one: []string{"doeeqiy", "yabhbqe", "twckqte"},
			},
			a: ans{
				3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, minDeletionSize(p.one), "input:%v", p)
	}
}
