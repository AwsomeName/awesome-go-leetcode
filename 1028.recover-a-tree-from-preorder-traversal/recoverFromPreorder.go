package problem1028

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getNum(s string)(int, int){
    var tmp string
    idx := len(s)-1
    for i := 0 ; i < len(s); i++{
        if s[i] != '-' {
            tmp = tmp+string(s[i])
        }else{
            idx = i
            break
        }
    }
    val,_ := strconv.Atoi(tmp)
    // fmt.Println(val,idx)
    return val, idx
}

func gets(s string)int{
    idx := 0
    for i:=0; i < len(s); i++{
        if s[i] == '-' {
            idx++
        }else{
            break
        }
    }
    // fmt.Println('idx:',idx)
    return idx
}

func recoverFromPreorder(S string) *TreeNode {
    var root TreeNode
    if len(S) < 1 {
        return &root
    }

    stk := make([]*TreeNode,0)
    hdr := 0

    val, idx := getNum(S)
    // S = S[idx:]
    root = TreeNode{val, nil, nil}

    stk = append(stk, &root)
    pre := 0

    if idx == 0 || idx == len(S)-1 {
        S = ""
    }else{
        S = S[idx:]
    }

    for len(S) > 0 {
        tmphdr := stk[hdr]

        prelen := gets(S)
        // fmt.Println(gets(S))
        // fmt.Println(prelen)
        // fmt.Println(S[gets(S):])
        S = S[gets(S):]
        // fmt.Println(getNum(S))
        tmpval,id :=getNum(S)
        fmt.Println(S[id:], tmpval, "S:",S,"id:",id,"len:",len(S))

        if id == 0 || id == len(S)-1 {
            S = ""
        }else{
            S = S[id:]
        }


        if prelen == pre { // new right
            newright := TreeNode{tmpval, nil, nil}
            hdr --
            tmphdr = stk[hdr]
            tmphdr.Right = &newright
            // fmt.Println("right:",tmphdr.Val)
            stk[hdr+1] = &newright
            hdr++
        }else if prelen < pre {
            fmt.Println("if prelen < pre")
            fmt.Println(hdr, prelen, pre)
            hdr = prelen - 1
            fmt.Println(hdr, prelen, pre)
            tmphdr = stk[hdr]
            fmt.Println(tmphdr.Val)
            newright := TreeNode{tmpval, nil, nil}
            tmphdr.Right = &newright
            stk[hdr+1] = &newright
            hdr++
            pre = prelen
        }else if prelen > pre { // new left
            newleft := TreeNode{tmpval, nil, nil}
            tmphdr.Left = &newleft
            // fmt.Println("left:",tmphdr.Val)
            // hdr++
            if len(stk) > hdr+1 {
                stk[hdr+1] = &newleft
                // hdr++

            }else{
                stk = append(stk, &newleft)
            }
            pre ++
            hdr++
        }

    }

    return &root
}


