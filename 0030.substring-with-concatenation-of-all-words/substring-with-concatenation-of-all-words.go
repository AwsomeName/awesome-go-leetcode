package problem0030

func findSubstring(s string, words []string) []int {
    //init
    res:=[]int{}

    sLen,wordsNum:=len(s),len(words)
    //special
    if sLen==0 || wordsNum==0{
        return res
    }
    wordLen:=len(words[0])
    remainNum:=make(map[string]int,wordsNum)
    count:=0

    for initialIndex := 0; initialIndex < wordLen; initialIndex++ {
        index:=initialIndex
        count=initRecord(words,remainNum)

        for index+wordsNum*wordLen<=sLen{
            word:=s[index+count*wordLen:index+(count+1)*wordLen]
            remainTimes,ok:=remainNum[word]
            switch {
            case !ok:
                index+=wordLen*(count+1)
                if count!=0{
                    count=initRecord(words,remainNum)
                }
            case remainTimes==0:
                index,count=moveIndex(index,wordLen,count,s,remainNum)
            default:
                remainNum[word]--
                count++
                if count == wordsNum {
                    res = append(res, index)
                    index,count=moveIndex(index,wordLen,count,s,remainNum)
                }
            }
        }
    }
    return res
}
func initRecord(words []string,remainNum map[string]int)int{
    for _,word := range words{
        remainNum[word]=0
    }
    for _,word := range words{
        remainNum[word]++
    }
    return 0
}
func moveIndex(index,wordLen,count int,s string,remainNum map[string]int)(int,int){
    remainNum[s[index:index+wordLen]]++
    count--
    index += wordLen
    return index,count
}
