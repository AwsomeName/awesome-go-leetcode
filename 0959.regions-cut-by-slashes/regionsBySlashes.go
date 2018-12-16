package problems0959

func regionsBySlashes(gird []string) int{
    n := len(gird)
    count := n*n*4
    f := make(map[int]int, count)
    for i:=0; i< count; i++{
        f[i] = i
    }

    var find func(x int) int
    find = func(x int) int {
        if (x!=f[x]){
            f[x] = find(f[x])
        }
        return f[x]
    }

    get := func(i,j,k int)int{
        return (i*n+j) * 4 + k
    }

    uni := func(x,y int){
        x,y = find(x),find(y)
        if (x != y){
            f[x] = y
            count --
        }
    }


    for i:=0; i<n; i++{
        for j:=0;j<n;j++{
            if i>0 {
                uni(get(i-1,j,2),get(i,j,0))
            }
            if j>0 {
                uni(get(i,j-1,1),get(i,j,3))
            }
            if gird[i][j] != '/' {
                uni(get(i,j,0), get(i,j,1))
                uni(get(i,j,2), get(i,j,3))
            }
            if gird[i][j] != '\\' {
                uni(get(i,j,0), get(i,j,3))
                uni(get(i,j,2), get(i,j,1))
            }
        }
    }

    return count
}
