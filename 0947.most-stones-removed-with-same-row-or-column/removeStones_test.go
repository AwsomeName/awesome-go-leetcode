package problem0947

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type para struct {
  one [][]int
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
      p: para {
        one: [][]int{{0,0},{0,1},{1,0},{1,2},{2,1},{2,2}},
      },
      a: ans{
        5,
      },
    },
    question{
      p: para {
        one: [][]int{{0,0},{0,2},{1,1},{2,0},{2,2}},
      },
      a: ans{
        3,
      },
    },
    question{
      p: para{
        one: [][]int{{0,1},{1,2},{1,3},{3,3},{2,3},{0,2}},
      },
      a: ans{
        5,
      },
    },
    question{
      p: para{
        one: [][]int{{0,0}},
      },
      a: ans{
        0,
      },
    },
  }

  for _, q := range qs {
    a, p := q.a, q.p
    ast.Equal(a.ans, removeStones(p.one), "input:%v", p)
  }
}
