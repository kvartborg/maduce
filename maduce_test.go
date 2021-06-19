package maduce

import (
	"testing"
)

func TestCollectionFrom(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	var sum int
	From(slice).Reduce(&sum, func(n, sum int) int {
		return n + sum
	})

	if sum != 15 {
		t.Errorf("expected %d; got %d", 15, sum)
	}
}

func TestCollection_Filter(t *testing.T) {
	c := Collection{1, 2, 3, 4, 5}

	filter := func(n int) bool {
		return n > 3
	}

	if result := len(c.Filter(filter)); result != 2 {
		t.Errorf("Failed, expected %d; got %d", 2, result)
	}
}

func BenchmarkCollection_Filter(b *testing.B) {
	c := Collection{1, 2, 3, 4, 5}
	filter := func(v, i int) bool {
		_ = i
		return v > 3
	}

	for i := 0; i < b.N; i++ {
		c.Filter(filter)
	}
}

func BenchmarkFilterTheGoWay(b *testing.B) {
	c := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		var result []int
		for _, item := range c {
			if item > 3 {
				result = append(result, item)
			}
		}
	}
}

func TestCollection_Reduce(t *testing.T) {
	c := Collection{1, 2, 3, 4, 5}

	reducer := func(n, sum, i int) int {
		_ = i
		return n + sum
	}

	var sum int
	c.Reduce(&sum, reducer)

	if sum != 15 {
		t.Errorf("expected sum of %d; got %d", 10, sum)
	}
}

func BenchmarkCollection_Reduce(b *testing.B) {
	c := Collection{1, 2, 3, 4, 5}

	reducer := func(n, sum int) int {
		return n + sum
	}

	for i := 0; i < b.N; i++ {
		var sum int
		c.Reduce(&sum, reducer)
	}
}

func BenchmarkReduceTheGoWay(b *testing.B) {
	c := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		var sum int
		for _, item := range c {
			sum += item
		}
	}
}

func TestCollection_Map(t *testing.T) {
	c := Collection{1, 2, 3, 4, 5}

	mapper := func(n int) int {
		return n * 10
	}

	c = c.Map(mapper)

	var sum int
	c.Reduce(&sum, func(n, sum int) int {
		return n + sum
	})

	if sum != 150 {
		t.Errorf("expected sum of %d; got %d", 150, sum)
	}
}

func BenchmarkCollection_Map(b *testing.B) {
	c := Collection{1, 2, 3, 4, 5}

	mapper := func(n int) int {
		return n * 10
	}

	for i := 0; i < b.N; i++ {
		c.Map(mapper)
	}
}

func BenchmarkMapTheGoWay(b *testing.B) {
	c := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		result := make([]int, len(c))
		for i, n := range c {
			result[i] = n * 10
		}
	}
}
