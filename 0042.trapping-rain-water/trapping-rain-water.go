package problem0042

//import "fmt"

func trap(height []int) int{
    res := 0
//    for i,x:= range height{
    for i:= 0; i< len(height);i++{
        x := height[i]
        //fmt.Println("start:",i,x)
        j:= i+1
        hgtest := j
        for  j<len(height) && height[j] < x{
            if height[j] > height[hgtest] { hgtest = j }
            j++
        }
        if j == i+ 1 { continue }
        if j < len(height) && height[j] > height[hgtest] { hgtest = j }
        hgts := height[i+1:j]
        //if j>=len(height) {fmt.Println("j over")} else {fmt.Println("find right:",j,height[j],"length:",j-i)}
        if j < len(height) && hgtest == j {
            //fmt.Println("normal:",hgts)
            for _,y:= range hgts {
                //fmt.Println("x:",x,"y:",y,"x-y:",x-y)
                res += (x-y)
            }
            i = j - 1
            //fmt.Println("next i:",i,"res:",res)
        }else if j == len(height) && hgtest == j{
            for _,y:= range hgts {
                res += (x-y)
            }
            i = j - 1
        }else if j == len(height) && hgtest < j {
            //fmt.Println("no more equal:", hgts[:hgtest-i])
            for _,y:= range hgts[:hgtest-i] {
                res += (height[hgtest]-y)
            }
            i = hgtest - 1
        }
    }
    return res
}
