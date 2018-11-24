package problem0026

func removeDuplicates(a []int) int {
    if len(a) < 2 { return 1}
    tmp := a[0]
    cnt := 1
    for i := 1; i<len(a) ; i++{
        if tmp != a[i] {
            cnt ++
            tmp = a[i]
        }
    }
	return cnt
}
