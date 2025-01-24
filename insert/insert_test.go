package insert

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	t.Run("input already sorted", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Sort(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("simple unsorted input", func(t *testing.T) {
		in := []int{3, 2, 4, 1, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Sort(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("reversed unsorted input", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		want := []int{1, 2, 3, 4, 5}
		got := Sort(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func BenchmarkInsertSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort([]int{5, 9, 52, 86, 0, 1, 22, 75, 89, 256, -1})
	}
}

func TestInsertSort2(t *testing.T) {
	t.Run("input already sorted", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Sort2(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("simple unsorted input", func(t *testing.T) {
		in := []int{3, 2, 4, 1, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Sort2(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("reversed unsorted input", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		want := []int{1, 2, 3, 4, 5}
		got := Sort2(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func BenchmarkInsertSort2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort2([]int{5, 9, 52, 86, 0, 1, 22, 75, 89, 256, -1})
	}
}

func TestLowest(t *testing.T) {
	got := lowest([]int{1, 2, 3}, 0)
	if got != 0 {
		t.Fatalf("index should be 0; was %v", got)
	}
	got = lowest([]int{3, 2, 1}, 0)
	if got != 2 {
		t.Fatalf("index should be 0; was %v", got)
	}
	got = lowest([]int{9, 9, 9, 1, 9, 9, 9}, 0)
	if got != 3 {
		t.Fatalf("index should be 0; was %v", got)
	}
	got = lowest([]int{9, 9, 9, 1, 9, 9, 9}, 3)
	if got != 3 {
		t.Fatalf("index should be 0; was %v", got)
	}
}

func TestInsertSort3(t *testing.T) {
	t.Run("input already sorted", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Sort3(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("simple unsorted input", func(t *testing.T) {
		in := []int{3, 2, 4, 1, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Sort3(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("reversed unsorted input", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		want := []int{1, 2, 3, 4, 5}
		got := Sort3(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func BenchmarkInsertSort3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort3([]int{5, 9, 52, 86, 0, 1, 22, 75, 89, 256, -1})
	}
}
