package problem0953

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []string
	two string
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
				one: []string{"hello", "leetcode"},
				two: "hlabcdefgijkmnopqrstuvwxyz",
			},
			a: ans{
				true,
			},
		},
		question{
			p: para{
				one: []string{"word", "world", "row"},
				two: "worldabcefghijkmnpqstuvxyz",
			},
			a: ans{
				false,
			},
		},
		question{
			p: para{
				one: []string{"apple", "app"},
				two: "abcdefghijklmnopqrstuvwxyz",
			},
			a: ans{
				false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.ans, isAlienSorted(p.one, p.two), "input:%v", p)
	}
}
