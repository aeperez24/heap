package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func comparator(a interface{}, b interface{}) bool {
	return a.(int) < b.(int)
}
func TestShouldReturnTrueWhenHeapIsEmpty(t *testing.T) {
	heap := NewHeap(comparator)
	assert.True(t, heap.IsEmpty(), "expected empty true")
}

func TestShouldReturnFalseWhenHeapIsNotEmpty(t *testing.T) {
	heap := NewHeap(comparator)
	heap.Put(1)
	assert.False(t, heap.IsEmpty(), "expected empty false")
}

func TestShouldReturnMinimalItemWhenPeak(t *testing.T) {
	heap := NewHeap(comparator)
	heap.Put(2)
	heap.Put(3)
	heap.Put(1)
	heap.Put(4)
	heap.Put(5)
	assert.Equal(t, 1, heap.Peak(), "expected minimal 1")
	heap.Remove()
	assert.Equal(t, 2, heap.Peak(), "expected minimal 2")
	heap.Remove()
	assert.Equal(t, 3, heap.Peak(), "expected minimal 3")
	heap.Remove()
	assert.Equal(t, 4, heap.Peak(), "expected minimal 4")
	heap.Remove()
	assert.Equal(t, 5, heap.Peak(), "expected minimal 5")

}

func TestShouldReturnMaximiunlItemWhenPeak(t *testing.T) {
	maximunComparator := func(a interface{}, b interface{}) bool {
		return !comparator(a, b)
	}
	heap := NewHeap(maximunComparator)
	heap.Put(2)
	heap.Put(3)
	heap.Put(1)
	heap.Put(4)
	heap.Put(5)
	assert.Equal(t, 5, heap.Peak(), "expected minimal 5")
	heap.Remove()
	assert.Equal(t, 4, heap.Peak(), "expected minimal 4")
	heap.Remove()
	assert.Equal(t, 3, heap.Peak(), "expected minimal 3")
	heap.Remove()
	assert.Equal(t, 2, heap.Peak(), "expected minimal 2")
	heap.Remove()
	assert.Equal(t, 1, heap.Peak(), "expected minimal 1")

}
