package mergesort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("single item returns itself", func(t *testing.T) {
		got := Sort([]int{1})
		want := []int{1}
		assertEqual(got, want, t)
	})
	t.Run("short list sort", func(t *testing.T) {
		example := []int{9, 0, 6, 3, 8, 5, 2, 3}
		got := Sort(example)
		want := []int{0, 2, 3, 3, 5, 6, 8, 9}
		assertEqual(got, want, t)
	})
	t.Run("reversed list sort", func(t *testing.T) {
		example := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		got := Sort(example)
		want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		assertEqual(got, want, t)
	})
	t.Run("empty list sort", func(t *testing.T) {
		example := make([]int, 0, 0)
		got := Sort(example)
		want := []int{}
		assertEqual(got, want, t)
	})
	t.Run("short list sort", func(t *testing.T) {
		example := []int{3, 5, 1, 2}
		got := Sort(example)
		want := []int{1, 2, 3, 5}
		assertEqual(got, want, t)
	})
}

func TestMerge(t *testing.T) {
	t.Run("equal len arrays", func(t *testing.T) {
		left := []int{1, 3, 5}
		right := []int{0, 2, 4}
		want := []int{0, 1, 2, 3, 4, 5}
		got := Merge(left, right)
		assertEqual(got, want, t)
	})
	t.Run("shorter left array", func(t *testing.T) {
		left := []int{1}
		right := []int{0, 2, 4}
		want := []int{0, 1, 2, 4}
		got := Merge(left, right)
		assertEqual(got, want, t)
	})
	t.Run("shorter right array", func(t *testing.T) {
		left := []int{3, 6, 9}
		right := []int{0}
		want := []int{0, 3, 6, 9}
		got := Merge(left, right)
		assertEqual(got, want, t)
	})
	t.Run("empty right array", func(t *testing.T) {
		left := []int{3, 6, 9}
		right := []int{}
		want := []int{3, 6, 9}
		got := Merge(left, right)
		assertEqual(got, want, t)
	})
}

func assertEqual(got, want []int, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func Benchmark_merge(t *testing.B) {
	sample := intSample(1000)
	for n := 0; n < t.N; n++ {
		Sort(sample)
	}
}

func benchmarkMerge(nums []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(nums)
	}
}

func BenchmarkMerge1k(b *testing.B) {
	benchmarkMerge(intSample(1000), b)
}
func BenchmarkMerge10k(b *testing.B) {
	benchmarkMerge(intSample(10000), b)
}
func BenchmarkMerge100k(b *testing.B) {
	benchmarkMerge(intSample(100000), b)
}
func BenchmarkMerge500k(b *testing.B) {
	benchmarkMerge(intSample(500000), b)
}

func intSample(n int) []int {
	ary := make([]int, 0, n)
	for range n {
		ary = append(ary, rand.Int())
	}
	return ary
}

func TestIntSample(t *testing.T) {
	sample := intSample(5)
	if len(sample) != 5 {
		t.Errorf("expected 5, got %d", len(sample))
	}
	sample = intSample(10)
	if len(sample) != 10 {
		t.Errorf("expected 10, got %d", len(sample))
	}

	s1 := intSample(10)
	s2 := intSample(10)
	if reflect.DeepEqual(s1, s2) {
		t.Error("samples should be different!")
	}
}
