package problem0015

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]int
}

func Test_Problem0015(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			ans{[][]int{
				[]int{-4, -2, 6},
				[]int{-4, 0, 4},
				[]int{-4, 1, 3},
				[]int{-4, 2, 2},
				[]int{-2, -2, 4},
				[]int{-2, 0, 2},
			}},
		},
		question{
			para{[]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}},
			ans{[][]int{
				[]int{-5, 1, 4},
				[]int{-4, 0, 4},
				[]int{-4, 1, 3},
				[]int{-2, -2, 4},
				[]int{-2, 1, 1},
				[]int{0, 0, 0},
			}},
		},

		question{
			para{[]int{1, -1, -1, 0}},
			ans{[][]int{
				[]int{-1, 0, 1},
			}},
		},

		question{
			para{[]int{-1, 0, 1, 2, 2, 2, 2, -1, -4}},
			ans{[][]int{
				[]int{-4, 2, 2},
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			}},
		},
		question{
			para{[]int{0, 0, 0, 0, 0}},
			ans{[][]int{
				[]int{0, 0, 0},
			}},
		},
		question{
			para{[]int{1, 1, -2}},
			ans{[][]int{
				[]int{-2, 1, 1},
			}},
		},
		question{
			para{[]int{0, 0, 0}},
			ans{[][]int{
				[]int{0, 0, 0},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, threeSum(p.one), "输入:%v", p)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
}
