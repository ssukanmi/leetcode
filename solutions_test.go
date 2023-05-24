package main

import (
	"reflect"
	"testing"

	"github.com/ssukanmi/leetcode/solutions"
)

// Test TwoSum solution
func TestTwoSum(t *testing.T) {
	input1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}
	actual1 := solutions.TwoSum(input1, target1)
	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Result was incorrect for TwoSum test1, got: %v, exepected: %v.", actual1, expected1)
	}
	input2 := []int{3, 2, 4}
	target2 := 6
	expected2 := []int{1, 2}
	actual2 := solutions.TwoSum(input2, target2)
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Result was incorrect for TwoSum test2, got: %v, exepected: %v.", actual2, expected2)
	}
	input3 := []int{3, 3}
	target3 := 6
	expected3 := []int{0, 1}
	actual3 := solutions.TwoSum(input3, target3)
	if !reflect.DeepEqual(expected3, actual3) {
		t.Errorf("Result was incorrect for TwoSum test3, got: %v, exepected: %v.", actual3, expected3)
	}
}

// Test IsPalindrome solution
func TestIsPalindrome(t *testing.T) {
	input1 := 121
	expected1 := true
	actual1 := solutions.IsPalindrome2(input1)
	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Result was incorrect for IsPalindrome test1, got: %t, exepected: %t.", actual1, expected1)
	}
	input2 := -121
	expected2 := false
	actual2 := solutions.IsPalindrome2(input2)
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Result was incorrect for IsPalindrome test2, got: %t, exepected: %t.", actual2, expected2)
	}
	input3 := 10
	expected3 := false
	actual3 := solutions.IsPalindrome2(input3)
	if !reflect.DeepEqual(expected3, actual3) {
		t.Errorf("Result was incorrect for IsPalindrome test3, got: %t, exepected: %t.", actual3, expected3)
	}
}
