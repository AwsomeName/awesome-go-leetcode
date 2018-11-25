package problem0946

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
    fmt.Println("-------------")
    var stack [1000]int
    if len(pushed) < len(popped) { return false}
    stack_index := -1
    pushed_index := 0
    for i:= range popped{
        fmt.Println("pre to pop: ",popped[i])
        if stack_index == -1 {
            if pushed_index == len(pushed) {
                return false
            }else {
                stack_index ++
                stack[stack_index] = pushed[pushed_index]
                fmt.Println("push: ", pushed[pushed_index])
                pushed_index ++
            }
        }
        if popped[i] == stack[stack_index]{
            stack_index--
            continue
        }else {
            for popped[i] != stack[stack_index]{
                if pushed_index >= len(pushed) { return false }
                stack_index ++
                stack[stack_index] = pushed[pushed_index]
                fmt.Println("push: ", pushed[pushed_index])
                pushed_index++
            }
            stack_index --
        }
    }
    return true
}
