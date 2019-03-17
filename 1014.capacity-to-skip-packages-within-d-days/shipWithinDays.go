package problem1014

func shipWithinDays(weights []int, D int) int {
    sum := 0
    max := 0 //math.MaxInt32
    for _,x := range weights {
        sum += x
        if max < x {
            max = x
        }
    }

    avg := sum / D
    if sum % D != 0 {
        avg ++
    }

    min := 0
    if avg < max {
        min = max
    }else{
        min = avg
    }


    for {
        fmt.Println(min)
        cnt := 0
        index := 0

        for index < len(weights) {
            tmp := min
            for tmp > 0 && index < len(weights){
                if tmp >= weights[index] {
                    tmp -= weights[index]
                    index ++
                }else {
                    break
                }
            }

            cnt ++
        }

        if cnt <= D {
            return min
        }else {
            min ++
        }
    }


    // return min
}
