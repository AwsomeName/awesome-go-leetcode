package problem0945

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type para struct {
  one []int
}

type ans struct {
  min int
}

type question struct {
  p para
  a ans
}

func Test_twoSum(t *testing.T) {
  ast := assert.New(t)

  qs := []question{
    question{
      p: para {
        []int{1,2,2},
      },
      a: ans{
        1,
      },
    },
    question{
      p: para{
        []int{3,2,1,2,1,7},
      },
      a: ans{
        6,
      },
    },
  }

  for _, q := range qs {
    a, p := q.a, q.p
    ast.Equal(a.min, minIncrementForUnique(p.one), "input:%v", p)
  }
}
