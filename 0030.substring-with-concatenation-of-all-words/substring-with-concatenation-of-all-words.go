package problem0030

import "fmt"

func findSubstring(s string, words []string) []int {
    fmt.Println("-------------")
    Map_exp := make(map[string]int, len(words))
    lenOfAllWords := 0
    nWords := make([]string,len(words))
    wd_cnt := 0
    ///record word show times
    for i:= range words{
        if Map_exp[words[i]] == 0 {
            nWords[wd_cnt] = words[i]
            wd_cnt ++
        }
        Map_exp[words[i]] ++
        lenOfAllWords += len(words[i])
    }
    nWords = nWords[0:wd_cnt]
    fmt.Println("nWords:",len(nWords),nWords)

    res := make([]int, len(s))
    res_cnt := 0
    for i:= 0; i < len(s) - lenOfAllWords + 1; i++{
        Map_act := make(map[string]int, wd_cnt)
        ///从每个位置开始，检查一遍是否符合。
        ///从每个位置i 开始，统计总长之内，字符出现次数
        str_to_match := s[i:i+lenOfAllWords]
        fmt.Println("str to match:",str_to_match)
        for j:=0; j< len(str_to_match); {
            tmp_match := false
            for k:= range nWords{
                if j+len(nWords[k]) > len(str_to_match) { continue }
                tmp := str_to_match[j:j+len(nWords[k])]
                fmt.Println("tmp:",tmp)
                if nWords[k] == tmp {
                    Map_act[nWords[k]] ++
                    tmp_match = true
                    j += len(nWords[k])
                    break
                }
            }
            if tmp_match == false { break }
        }
        ///检查实际次数和真实次数是否相等
        match := false
        for j:= range nWords{
            if Map_exp[nWords[j]] == Map_act[nWords[j]] {
                match = true
                continue
            }
            match = false
            break
        }
        fmt.Println("Map_act:",Map_act)
        fmt.Println("Map_exp:",Map_exp)
        if match {
            res[res_cnt] = i
            res_cnt ++
        }
    }
    res = res[0:res_cnt]
    fmt.Println("res:",res)
    return res
}
