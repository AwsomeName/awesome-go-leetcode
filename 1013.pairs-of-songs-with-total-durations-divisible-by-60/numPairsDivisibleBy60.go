package problem1013

func numPairsDivisibleBy60(time []int) int {
    cnt := make([]int, 61)


    for i := range time {
        if time[i] > 0 && time[i] % 60 == 0 {
            cnt[60]++
        }else{
            time[i] = time[i] % 60
            cnt[time[i]]++
        }
    }

    ans := 0

    ans += (cnt[60] * (cnt[60]-1) /2)

    for i,x := range cnt {
        j := 60 - i

        if i == 30  {
            ans += (x * (x-1) /2)
        }else{
            ans += (x * cnt[j])
        }
        cnt[i] = 0
        cnt[j] = 0
    }

    // ans += sum(cnt[60])
    return ans
}
