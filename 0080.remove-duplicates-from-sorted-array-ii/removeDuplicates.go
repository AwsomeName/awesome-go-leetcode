package problem0080

// func removeDuplicates(nums []int) int {
//     if len(nums)==0 {
//         return 0
//     }
//     if len(nums)<=2 {
//         return len(nums)
//     }
//     cnt := 0
//     ans := 2

//     for i:=2; i< len(nums); i++{
//         if nums[i]==nums[i-cnt-2] {
//             cnt ++
//         }else{
//             ans ++
//         }
//         nums[i-cnt] = nums[i]
//     }

//     return ans
// }

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}

	// 大于2直接移除
	// 把后往前移

	rmFlag := 1
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == tmp {
			rmFlag++
			if rmFlag >= 3 {
				//fmt.Println(i, nums[i])
				//for j:=i;j<len(nums);j++ {
				//
				//}
				//fmt.Println(nums[0:i])
				//fmt.Println(nums[i+1:])

				nums = append(nums[0:i], nums[i+1:]...)
				i--
				//fmt.Println(nums)
			}
		} else {
			tmp = nums[i]
			rmFlag = 1
		}
	}
	return len(nums)
}
