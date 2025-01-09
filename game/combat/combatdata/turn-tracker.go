package combatdata

import (
	"container/heap"
	"quasar/characters"
)

// A PriorityQueue implements heap.Interface and holds Characters.
type PriorityQueue []characters.Character

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].GetTurnmeter() > pq[j].GetTurnmeter() }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(characters.Character))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type TurnTracker struct {
	Queue PriorityQueue
}

func NewTurnTracker(c []characters.Character) *TurnTracker {
	pq := make(PriorityQueue, len(c))
	for _, char := range c {
		pq.Push(char)
	}
	heap.Init(&pq)
	return &TurnTracker{Queue: pq}
}

func (tt *TurnTracker) Tick() characters.Character {
	for i := 0; i < len(tt.Queue); i++ {
		tt.Queue[i].Tick()
	}
	// Re-heapify to maintain priority
	heap.Init(&tt.Queue)

	// Check if the top character is ready for a turn
	top := tt.Queue[0]
	if top.GetTurnmeter() >= 1000 {
		return top
	}
	return nil
}
