package problem0948

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type para struct {
  one []int
  two int
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
        one: []int{100},
        two: 50,
      },
      a: ans{
        0,
      },
    },
    question{
      p: para {
        one: []int{100,200},
        two: 150,
      },
      a: ans{
        1,
      },
    },
    question{
      p: para{
        one: []int{100,200,300,400},
        two: 200,
      },
      a: ans{
        2,
      },
    },
  }

  for _, q := range qs {
    a, p := q.a, q.p
    ast.Equal(a.ans, bagOfTokensScore(p.one, p.two), "input:%v", p)
  }
}
