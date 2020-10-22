package utils

import (
	"container/heap"

	"gocv.io/x/gocv"
)

// Frame contains the priority and image to be written into a video later
type Frame struct {
	Image    gocv.Mat
	Priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Frame

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes x into pq
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	frame := x.(*Frame)
	frame.index = n
	*pq = append(*pq, frame)
}

// Pop returns frame with minimum priority
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	frame := old[n-1]
	old[n-1] = nil   // avoid memory leak
	frame.index = -1 // for safety
	*pq = old[0 : n-1]
	return frame
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(frame *Frame, priority int) {
	frame.Priority = priority
	heap.Fix(pq, frame.index)
}
