package problem0949

import "sort"

func largestTimeFromDigits(A []int) string{
	sort.Ints(A)
    if A[0] == A[3] {
        switch {
            case A[0] == 0 :
                return "00:00"
            case A[0] == 1 :
                return "11:11"
            case A[0] == 2 :
                return "22:22"
            default:
                return ""
        }
    }
    maxTime := []int{0,0}
    for i:= range A{
        for j:= range A{
            if j==i {continue}
            for m:= range A{
                if m==i || m==j {continue}
                for n:= range A{
                    if n==m||n==i||n==j {
                        continue
                    }
                    hour := A[i]*10 + A[j]
                    min := A[m]*10 + A[n]
                    if hour > 23 || min > 59 {
                        continue
                    }
                    if hour > maxTime[0] || min > maxTime[1] {
                        maxTime = []int{hour, min}
                    }
                }
            }
        }
    }
    if maxTime[0] == 0 && maxTime[1] == 0 {
       return ""
    }
    var strMaxTime string
    strMaxTime = strMaxTime + string(maxTime[0]/10+'0')
    strMaxTime = strMaxTime + string(maxTime[0]%10+'0')
    strMaxTime = strMaxTime + string(':')
    strMaxTime = strMaxTime + string(maxTime[1]/10+'0')
    strMaxTime = strMaxTime + string(maxTime[1]%10+'0')
	return strMaxTime
}
