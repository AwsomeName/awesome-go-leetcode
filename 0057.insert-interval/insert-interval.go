package problem0057

// Interval is a couple of int
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	a := intervals

	lenA := len(a)

	if lenA == 0 {
		return []Interval{newInterval}
	}

	if newInterval.End < a[0].Start {
		return append([]Interval{newInterval}, a...)
	}

	if newInterval.Start > a[lenA-1].End {
		return append(a, newInterval)
	}

	res := make([]Interval, 0, len(a))
	for i := range a {
		if isOverlap(a[i], newInterval) {
			newInterval = merge(a[i], newInterval)
			if i == lenA-1 {
				res = append(res, newInterval)
			}
			continue
		}

		if a[i].End < newInterval.Start {
			res = append(res, a[i])
			continue
		}

		if newInterval.End < a[i].Start {
			res = append(res, newInterval)
			res = append(res, a[i:]...)
			break
		}
	}
	return res
}

func isOverlap(a, b Interval) bool {
	return (a.Start <= b.Start && b.Start <= a.End) || (a.Start <= b.End && b.End <= a.End) || (b.Start <= a.Start && a.Start <= b.End)
}

func merge(a, b Interval) Interval {
	return Interval{
		Start: min(a.Start, b.Start),
		End:   max(a.End, b.End),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
