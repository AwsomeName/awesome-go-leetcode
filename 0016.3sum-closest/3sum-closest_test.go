package problem0016

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
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0016(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{-1, 2, 1, -4}, 1},
			ans{2},
		},
		question{
			para{[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, 1},
			ans{0},
		},
		question{
			para{[]int{1,6,9,14,16,70}, 81},
			ans{80},
		},
		question{
			para{[]int{-1,0,1,1,55}, 3},
			ans{2},
		},
        question{
            para{[]int{6,-18,-20,-7,-15,9,18,10,1,-20,-17,-19,-3,-5,-19,10,6,-11,1,-17,-15,6,17,-18,-3,16,19,-20,-3,-17,-15,-3,12,1,-9,4,1,12,-2,14,4,-4,19,-20,6,0,-19,18,14,1,-15,-5,14,12,-4,0,-10,6,6,-6,20,-8,-6,5,0,3,10,7,-2,17,20,12,19,-13,-1,10,-1,14,0,7,-3,10,14,14,11,0,-4,-15,-8,3,2,-5,9,10,16,-4,-3,-9,-8,-14,10,6,2,-12,-7,-16,-6,10},-52},
            ans{-52},
        },
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, threeSumClosest(p.one, p.two), "输入:%v", p)
	}
}
