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
        if len(res) < 1 {
            res = append(res,deck[i])
            continue
        }

        if len(res) > 1{
            tres := make([]int, 2)
            tres[0] = deck[i]
            tres[1] = res[len(res)-1]
            for j:=0; j<len(res)-1; j++{
               tres = append(tres,res[j])
            }
            res = tres
        }else {
            tres := make([]int,2)
            tres[0] = deck[i]
            tres[1] = res[0]
            res = tres
        }
    }
    return res
}
