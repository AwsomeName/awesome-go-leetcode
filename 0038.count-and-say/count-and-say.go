package problem0038

//import "fmt"

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	res := "11"
	for i := 3; i <= n && i <= 30; i++ {
		//        fmt.Println("before  i:",i,"res:",res)
		tmp := ""
		cnt := 0
		for j := 0; j < len(res); j++ {
			if j+1 < len(res) {
				if res[j] == res[j+1] {
					cnt++
				} else {
					cnt++
					tmp = tmp + string(cnt+'0') + string(res[j])
					cnt = 0
				}
			} else {
				cnt++
				tmp = tmp + string(cnt+'0') + string(res[j])
			}
		}
		res = tmp
		//        fmt.Println("after   i:",i,"res:",res)
	}
	return res
}
