package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var queque = new(MaxQueue)
	for _, n := range nums {
		if queque.Size() < k {
			queque.PushBack(n)
			continue
		}
		queque.println()
		result = append(result, queque.Max())
		queque.PopFront()
		queque.PushBack(n)
	}
	if queque.Size() == k {
		result = append(result, queque.Max())
	}
	return result
}

type MaxQueue struct {
	queque []int
	data   []int
}

func (queue *MaxQueue) PushBack(n int) {
	queue.data = append(queue.data, n)
	for {
		if len(queue.queque) == 0 {
			break
		}
		if queue.queque[len(queue.queque)-1] >= n {
			break
		}
		queue.queque = queue.queque[:len(queue.queque)-1]
	}
	queue.queque = append(queue.queque, n)
}

func (queue *MaxQueue) PopFront() int {
	var first = queue.data[0]
	queue.data = queue.data[1:]
	if queue.queque[0] == first {
		queue.queque = queue.queque[1:]
	}
	return first
}

func (queue *MaxQueue) println() {
	fmt.Printf("%v\n", queue.data)
}
func (queque *MaxQueue) Max() int {
	return queque.queque[0]
}

func (queque *MaxQueue) Size() int {
	return len(queque.data)
}

func main() {
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}
