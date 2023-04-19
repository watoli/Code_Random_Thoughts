package main

import "testing"

func TestSearch(t *testing.T) {
	// Test case 1
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	expected1 := 4
	result1 := search(nums1, target1)
	if result1 != expected1 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums1, target1, result1, expected1)
	}

	// Test case 2
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	expected2 := -1
	result2 := search(nums2, target2)
	if result2 != expected2 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums2, target2, result2, expected2)
	}

	// Test case 3
	nums3 := []int{1}
	target3 := 1
	expected3 := 0
	result3 := search(nums3, target3)
	if result3 != expected3 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums3, target3, result3, expected3)
	}

	// Test case 4
	nums4 := []int{5, 7, 8, 9, 10}
	target4 := 6
	expected4 := -1
	result4 := search(nums4, target4)
	if result4 != expected4 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums4, target4, result4, expected4)
	}
}
