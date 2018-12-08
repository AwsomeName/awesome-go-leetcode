package problem0043

//import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, 112*112)
	start := 0
	index := 0
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			start = len(num1) - i + len(num2) - j - 2
			xx := int(num1[i] - '0')
			yy := int(num2[j] - '0')
			tmp := xx * yy
			index = 0
			for tmp > 0 {
				//fmt.Println("index:",start+index,"Len:",len(res),"i:",i,"j:",j,"index:",index)
				//fmt.Println("res:",res[start+index],"tmp:",tmp,"new:",res[start+index]+(tmp%10))
				tmp += res[start+index]
				res[start+index] = tmp % 10
				index++
				tmp = tmp / 10
				//fmt.Println("res:",res[:start+index])
			}
		}
	}
	var res_str string
	for x := start + index - 1; x >= 0; x-- {
		res_str += string(res[x] + '0')
	}
	return res_str
}
