
type MinStack struct {
	a [][2]int
}

func Constructor() MinStack {
	return MinStack{make([][2]int, 0)}
}

func (this *MinStack) Push(val int) {
	if len(this.a) > 0 {
		a := min(val, this.a[len(this.a)-1][1])
		this.a = append(this.a, [2]int{val, a})
	} else {
		this.a = append(this.a, [2]int{val, val})
	}
}

func (this *MinStack) Pop() {
	this.a = this.a[:len(this.a)-1]
}

func (this *MinStack) Top() int {
	return this.a[len(this.a)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.a[len(this.a)-1][1]
}
