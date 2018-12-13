package problem0046

func permute(nums []int) [][]int {
    n := len(nums)
    var ans [][]int
    taken := make([]bool,n)
    vector := make([]int, n)
    makePermutation(0,n,nums,vector,taken,&ans)
    return ans
}

func makePermutation(cur, n int, nums,vector []int, taken []bool, ans *[][]int){
    if cur == n {
        tmp := make ([]int,n)
        copy (tmp, vector)
        *ans = append(*ans,tmp)
        return
    }

    for i := 0; i<n; i++ {
        if !taken[i] {
            taken[i] = true
            vector[cur] = nums[i]
            makePermutation(cur+1, n, nums, vector, taken, ans)
            taken[i] = false
        }
    }
}
