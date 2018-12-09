package problem0954

//import "fmt"
import "sort"

func canReorderDoubled(A []int) bool {
	//    fmt.Println("-------------------")
	sort.Ints(A)

	index := 0
	left, right := 0, len(A)-1
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if A[mid] == 0 {
			index = mid
			break
		} else if A[mid] < 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//   fmt.Println("mid:",mid,left,right)
	if A[index] == 0 {
		if index%2 == 1 {
			if A[index-1] != 0 {
				return false
			} else {
				index = index - 1
			}
		} else {
			if A[index+1] != 0 {
				return false
			}
		}
	} else {
		if mid%2 == 1 {
			index = mid + 1
		} else {
			index = mid
		}
	}
	//    fmt.Println("index:",index)
	PA := A[index:]
	NA := A[:index]
	// fmt.Println("PA:",PA)
	//  fmt.Println("NA:",NA)

	for i := 0; i < len(PA)/2; i++ {
		if PA[i]*2 != PA[len(PA)/2+i] {
			//           fmt.Println("cmp:",PA[i],PA[len(PA)/2+i])
			return false
		}
	}
	for i := 0; i < len(NA)/2; i++ {
		if NA[i] != NA[len(NA)/2+i]*2 {
			//            fmt.Println("cmp:",NA[i],"1x:",NA[len(NA)/2+i], "2x:",2*NA[len(NA)/2+i])
			return false
		}
	}
	return true
}
