package problem1028

func baseNeg2(N int) string {
    if N==0 {
        return "0"
    }
    if N == 1 {
        return "1"
    }
    var tmp string
    if N & 1 == 1 {
        tmp = "1"
    }else{
        tmp = "0"
    }
    return baseNeg2(-(N>>1)) + tmp
}
