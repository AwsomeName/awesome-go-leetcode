package problem0950

import "sort"
//import "fmt"

func deckRevealedIncreasing(deck []int) []int {
    if len(deck) < 1 {
        return nil
    }
    if len(deck) == 1 {
        return deck
    }
    //res := make([]int, len(deck))
    var res []int
	sort.Ints(deck)
    if len(deck)==2 {
        return deck
    }
    for i:= len(deck)-1; i >=0; i--{
        tmp:= res
        if len(tmp) < 1 {
            res = append(res,deck[i])
            continue
        }

        if len(tmp) > 1{
            tres := make([]int, 2)
            tres[0] = deck[i]
            tres[1] = tmp[len(tmp)-1]
            for j:=0; j<len(tmp)-1; j++{
               tres = append(tres,tmp[j])
            }
            res = tres
        }else {
            tres := make([]int,2)
            tres[0] = deck[i]
            tres[1] = tmp[0]
            res = tres
        }
    }
    return res
}
