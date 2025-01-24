package insert

import (
// "fmt"
)

/*
Insert sort works by scanning thru the items and placing the current
item in its ordered location to the left of the pointer.
*/
func Sort(nums []int) []int {
	//	fmt.Println("start:", nums)
	for j := 1; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[j] < nums[i] { // j goes before i
				item := nums[j]
				shiftright(nums, i, j)
				nums[i] = item
				break
			}
		}
		//		fmt.Println(nums)
	}

	return nums
}

/*
inserts the item at [j] before the item[i].  To do this, we
need to shift i-j right to make room to insert j.  We can
make room by pulling j out and shifting right from i into j.
*/
func shiftright(nums []int, i, j int) {
	if !(i < len(nums)-1) {
		panic("cannot shift last item")
	}
	// shift  starting from the right end
	for p := j; p > i; p-- {
		nums[p] = nums[p-1]
	}
}

// this version of sort does not shift-right, but instead swaps
// the larger item with the smaller one.
func Sort2(nums []int) []int {
	for j := 1; j < len(nums); j++ {
		for i := 0; i < len(nums); i++ {
			if nums[i] > nums[j] {
				swap(nums, i, j)
			}
		}
	}
	return nums
}

func swap(nums []int, i, j int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}

// this version scans for the lowest and inserts it into the
// current lowest slot.
func Sort3(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		lx := lowest(nums, i)
		if lx != i {
			swap(nums, i, lx)
		}
	}
	return nums
}

// returns the index of the lowest in the splice,
// starting at the ith element.
func lowest(nums []int, start int) int {
	min := start
	for i := start; i < len(nums); i++ {
		if nums[i] < nums[min] {
			min = i
		}
	}
	return min
}
