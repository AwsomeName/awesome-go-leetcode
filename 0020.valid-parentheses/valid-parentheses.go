package problem0020

func isValid(s string) bool {
    if s == "" {
        return true
    }
    stack :=  []rune{}
    for _,v := range s {
        if len(stack) == 0 {
            stack = append(stack, v)
            continue
        }
        // '(' is 40 and ')' is 41 , sum = 81
        res := int(v) + int(stack[len(stack) - 1])
        if res == 184 || res == 248 || res == 81 {
            stack = stack[:len(stack) - 1]// out
        }else {
            stack = append(stack, v)// in
        }
    }
    if len(stack) == 0{
        return true
    }
    return false
}
