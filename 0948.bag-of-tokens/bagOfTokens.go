package problem0948

import "sort"

func bagOfTokensScore(tokens []int, P int) int {
    sort.Ints(tokens)
    score := 0
    max_score := 0
    for len(tokens) > 0 {
        if P < tokens[0] && score == 0 { break }
        ///首先用P换取最小的token以得分
        for len(tokens)> 0 && P >= tokens[0] {
            if len(tokens)<= 1 {
                if max_score <  score + 1 {
                    return score + 1
                }else {
                    return max_score
                }
            }else {
                P -= tokens[0]
                score ++
                tokens = tokens[1:len(tokens)]
                if max_score <  score {
                    max_score = score
                }
            }
        }
        ///用分数换取最大的token以加分
        if score > 0 && len(tokens)>0 {
            score --
            P += tokens[len(tokens)-1]
            tokens = tokens[0:len(tokens)-1]
        }
    }
    return max_score
}
