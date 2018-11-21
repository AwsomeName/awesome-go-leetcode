package problem0020

func isValid(str string) bool {
    var stack [100]byte
    var head int
    if str == "" { return true }
    for i := range str{
        switch(str[i]){
            case '(':
                stack[head]='('
                head ++
            case '[':
                stack[head]='['
                head ++
            case '{':
                stack[head]='{'
                head ++
            case ')':
                if head == 0 {return false}
                if stack[head-1] == '(' { head -- }else { return false }
            case ']':
                if head == 0 {return false}
                if stack[head-1] == '[' { head -- }else { return false }
            case '}':
                if head == 0 {return false}
                if stack[head-1] == '{' { head -- }else { return false }
            default:
                return false
        }
    }
    if head != 0 { return false }
    return true
}
