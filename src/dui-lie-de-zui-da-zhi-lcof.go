package main

type MaxQueue struct {
	queque []int
	data   []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.queque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	for {
		if len(this.queque) == 0 {
			break
		}
		if this.queque[len(this.queque)-1] >= value {
			break
		}
		this.queque = this.queque[:len(this.queque)-1]
	}
	this.queque = append(this.queque, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	var first = this.data[0]
	if this.queque[0] == this.data[0] {
		this.queque = this.queque[1:]
	}
	this.data = this.data[1:]
	return first
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
