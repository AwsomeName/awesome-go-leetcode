package problem0961

func repeatedNTimes(A []int) int {
    marked:=make(map[int]bool)
    for _,a := range A {
        if marked[a] {
            return a
        }else{
            marked[a] = true
        }
    }
    return 0
}
