package problem0028

func strStr(haystack string, needle string) int {
    if needle == "" { return 0 }
    var match bool
    for i:=0; i <len(haystack)-len(needle)+1; i++{
        for j:= range needle{
            if haystack[i+j] != needle[j]{
                match = false
                break
            }else {
                match = true
            }
        }
        if match  { return i}
    }

	return -1
}
