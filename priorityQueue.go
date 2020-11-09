package priorityQueue

import (
	"container/heap"
	"log"
	"sync"
)

// An Item is something we manage in a priority queue.
type Item struct {
	Value    int // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

var lock sync.RWMutex

func (pq PriorityQueue) Len() int {
	lock.Lock()
	defer lock.Unlock()
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	lock.Lock()
	defer lock.Unlock()
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	lock.Lock()
	defer lock.Unlock()
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (pq *PriorityQueue) Peek() interface{} {
	lock.Lock()
	defer lock.Unlock()
	p := *pq
	item := p[0]
	return item
}
func (pq *PriorityQueue) Push(x interface{}) {
	lock.Lock()
	defer lock.Unlock()
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
	log.Println(pq)
}

func (pq *PriorityQueue) Pop() interface{} {
	lock.Lock()
	defer lock.Unlock()
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	lock.Lock()
	defer lock.Unlock()
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
