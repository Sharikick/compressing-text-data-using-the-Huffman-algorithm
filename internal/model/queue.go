package model

import (
	"log/slog"
	"os"
)

type Node struct {
	Left, Right *Node
	Frequency   int
	Char        rune
}

type PriorityQueue struct {
	items []*Node
	log   *slog.Logger
}

func CreatePriorityQueue(log slog.Logger) *PriorityQueue {
	return &PriorityQueue{
		items: make([]*Node, 0, 40),
		log:   &log,
	}
}

func (queue *PriorityQueue) Length() int {
	return len(queue.items)
}

func (queue *PriorityQueue) Swap(i, j int) {
	queue.items[i], queue.items[j] = queue.items[j], queue.items[i]
}

func (queue *PriorityQueue) siftUp(index int) {
	parent := (index - 1) / 2
	for parent >= 0 {
		if queue.items[parent].Frequency > queue.items[index].Frequency {
			queue.Swap(index, parent)
		} else {
			break
		}
		index = parent
		parent = (index - 1) / 2
	}
}

func (queue *PriorityQueue) siftDown(index int) {
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2

	for leftIndex <= queue.Length() {
		swapIndex := leftIndex
		if rightIndex <= queue.Length() && queue.items[swapIndex].Frequency > queue.items[rightIndex].Frequency {
			swapIndex = rightIndex
		}
		if queue.items[index].Frequency > queue.items[swapIndex].Frequency {
			queue.Swap(index, swapIndex)
			index = swapIndex
		} else {
			break
		}
	}
}

func (queue *PriorityQueue) Add(node *Node) {
	queue.items = append(queue.items, node)
	queue.siftUp(queue.Length() - 1)
}

func (queue *PriorityQueue) PopRoot() *Node {
	if queue.Length() == 0 {
		queue.log.Error("The list is empty")
		os.Exit(1)
	}

	node := queue.items[0]
	lastIndex := queue.Length() - 1
	queue.items[0] = queue.items[lastIndex]
	queue.items = queue.items[:lastIndex]
	queue.siftDown(0)
	return node
}
