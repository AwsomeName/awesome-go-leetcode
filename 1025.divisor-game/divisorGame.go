package problem1025

func divisorGame(N int) bool {
    if N == 1{
        return false
    }
    if N == 2 {
        return true
    }

    record := make([]bool , N)

    record[0], record[1], record[2] = false , true, false//1,2,3

    for i:=3; i<N; i++{
        // N = i+1
        n := i+1
        for j := 1; j < n; j++{
            if (n % j == 0) && !record[i-j] {
                record[i] = true
                break
            }
        }
    }

    return record[N-1]
}
