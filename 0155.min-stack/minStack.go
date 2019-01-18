package problem0155

type MinStack struct {
	Vals    []int
	minVals []int
	hdr     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	// var
	// Val = make([]int, 0)
	// minVal = make([]int,0)
	// hdr = 0
	return MinStack{[]int{}, []int{}, 0}
}

func (this *MinStack) Push(x int) {
	if this.hdr == 0 {
		if len(this.Vals) == 0 {
			this.Vals = append(this.Vals, x)
			this.minVals = append(this.minVals, x)
		} else {
			this.Vals[this.hdr] = x
			this.minVals[this.hdr] = x
		}
	} else {
		if len(this.Vals) <= this.hdr {
			this.Vals = append(this.Vals, x)
			if x < this.minVals[this.hdr-1] {
				this.minVals = append(this.minVals, x)
			} else {
				this.minVals = append(this.minVals, this.minVals[this.hdr-1])
			}
		} else {
			this.Vals[this.hdr] = x
			if x < this.minVals[this.hdr-1] {
				this.minVals[this.hdr] = x
			} else {
				this.minVals[this.hdr] = this.minVals[this.hdr-1]
			}
		}
	}
	this.hdr++
}

func (this *MinStack) Pop() {
	this.hdr--
}

func (this *MinStack) Top() int {
	return this.Vals[this.hdr-1]
}

func (this *MinStack) GetMin() int {
	return this.minVals[this.hdr-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
