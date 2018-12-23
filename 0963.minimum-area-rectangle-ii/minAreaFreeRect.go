package problem0963

func minAreaFreeRect(points [][]int) float64 {
	if len(points) < 4 {
		return 0
	}
	mid := make(map[string][][]int)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dis := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			cx := float32(points[j][0]+points[i][0]) / 2
			cy := float32(points[j][1]+points[i][1]) / 2
			key := "" + string(dis) + "+" + strconv.FormatFloat(float64(cx), 'E', -1, 32) + "+" + strconv.FormatFloat(float64(cy), 'E', -1, 32)
			if len(mid[key]) == 0 {
				mid[key] = make([][]int, 0)
			}
			mid[key] = append(mid[key], []int{i, j})
		}
	}

	res := math.MaxFloat64
	for _, set := range mid {
		if len(set) > 1 {
			for i := 0; i < len(set); i++ {
				for j := i + 1; j < len(set); j++ {
					p1 := set[i][0]
					p2 := set[j][0]
					p3 := set[j][1]
					tmp := math.Sqrt(float64(points[p1][0]-points[p2][0])*float64(points[p1][0]-points[p2][0])+float64(points[p1][1]-points[p2][1])*float64(points[p1][1]-points[p2][1])) * math.Sqrt(float64(points[p1][0]-points[p3][0])*float64(points[p1][0]-points[p3][0])+float64(points[p1][1]-points[p3][1])*float64(points[p1][1]-points[p3][1]))
					res = min(res, tmp)
				}
			}
		}
	}

	if res < math.MaxFloat64 {
		return res
	} else {
		return 0.0
	}
}
func min(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}

}
