package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	type Student struct {
		Name string
		Age  int
	}

	nums1 := []int{1, 4, 2, 5, 3}
	nums2 := []int{1, 4, 2, 5, 3}
	nums3 := []int{1, 4, 2, 5, 3}
	sort.Ints(nums1)
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	slices.Sort(nums3)
	slices.SortFunc(nums3, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	students := []Student{{Name: "A", Age: 10}, {Name: "B", Age: 8}, {Name: "C", Age: 9}}
	slices.SortFunc(students, func(a, b Student) int {
		return a.Age - b.Age
	})

	// Output:
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(nums3)
	fmt.Println(students)
}
