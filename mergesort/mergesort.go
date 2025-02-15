package mergesort

import (
// "fmt"
)

/*
Sort takes a slice of ints and returns the slice sorted using the
merge sort algorith as follows:

if len([]) == 1 then done
else recurse:
sort[1..n/2] and sort[n/2+1..n]
merge s1 + s2

merge subroutine:
20  12
13  11
7  9
2  1  <- head of list

take the lowest of two heads() into output list.
repeat until both lists are empty.\

*/

func Sort(nums []int) []int {
	//fmt.Printf("Sort: %v \n", nums)
	switch len(nums) {
	case 0:
		return nums
	case 1:
		return nums
	default:
		la := nums[:len(nums)/2]
		ra := nums[len(nums)/2:]
		la = Sort(la)
		ra = Sort(ra)
		return Merge(la, ra)
	}
}

func Merge(la, ra []int) []int {
	//fmt.Printf("merge l=%v, r=%v \n", la, ra)
	result := make([]int, 0, len(la)+len(ra))
	var l, r int

	for {
		switch {
		case l == len(la) && r == len(ra):
			return result
		case l == len(la):
			result = append(result, ra[r])
			r++
		case r == len(ra):
			result = append(result, la[l])
			l++
		case ra[r] < la[l]:
			result = append(result, ra[r])
			r++
		default:
			result = append(result, la[l])
			l++
		}
	}
}
