package heap

type heapImpl struct {
	nodes   []interface{}
	compare Comparator
}

type Heap interface {
	Put(interface{})
	Remove()
	Peak() interface{}
	IsEmpty() bool
}

func NewHeap(comparator Comparator) Heap {
	return &heapImpl{
		nodes: make([]interface{}, 0), compare: comparator,
	}
}

type Comparator func(interface{}, interface{}) bool

func (heap heapImpl) IsEmpty() bool {
	return len(heap.nodes) == 0
}
func (heap *heapImpl) Put(node interface{}) {
	heap.nodes = append(heap.nodes, node)
	heap.sortUp(len(heap.nodes) - 1)
}
func (heap *heapImpl) Remove() {
	heap.swap(0, len(heap.nodes)-1)
	heap.nodes = heap.nodes[:len(heap.nodes)-1]
	heap.sortDown(0)
}
func (heap heapImpl) Peak() interface{} {
	return heap.nodes[0]
}

func (heap *heapImpl) swap(pos1 int, pos2 int) {
	aux := heap.nodes[pos1]
	heap.nodes[pos1] = heap.nodes[pos2]
	heap.nodes[pos2] = aux
}

func (heap *heapImpl) sortDown(pos int) {

	leftSonIndex := 2*pos + 1
	rightSonIndex := 2*pos + 2
	if leftSonIndex >= len(heap.nodes) {
		return
	}

	selectedSon := heap.compareIndex(leftSonIndex, rightSonIndex)
	toBeSorted := heap.compare(heap.nodes[selectedSon], heap.nodes[pos])
	if toBeSorted {
		heap.swap(selectedSon, pos)
		heap.sortDown(selectedSon)
	}
}

func (heap *heapImpl) sortUp(pos int) {
	if pos <= 0 {
		return
	}
	parentPos := (pos - 1) / 2
	toBeSorted := heap.compare(heap.nodes[pos], heap.nodes[parentPos])
	if toBeSorted {
		heap.swap(pos, parentPos)
		heap.sortUp(parentPos)
	}
}

func (heap heapImpl) compareIndex(leftSonIndex int, rightsonIndex int) int {
	if rightsonIndex >= len(heap.nodes) {
		return leftSonIndex
	}
	if heap.compare(heap.nodes[leftSonIndex], heap.nodes[rightsonIndex]) {
		return leftSonIndex
	}
	return rightsonIndex
}
