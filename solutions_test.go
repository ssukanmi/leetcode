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

// Test RomanToInt solution
func TestRomanToInt(t *testing.T) {
	input1 := "III"
	expected1 := 3
	actual1 := solutions.RomanToInt(input1)
	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Result was incorrect for RomanToInt test1, got: %d, exepected: %d.", actual1, expected1)
	}
	input2 := "LVIII"
	expected2 := 58
	actual2 := solutions.RomanToInt(input2)
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Result was incorrect for RomanToInt test2, got: %d, exepected: %d.", actual2, expected2)
	}
	input3 := "MCMXCIV"
	expected3 := 1994
	actual3 := solutions.RomanToInt(input3)
	if !reflect.DeepEqual(expected3, actual3) {
		t.Errorf("Result was incorrect for RomanToInt test3, got: %d, exepected: %d.", actual3, expected3)
	}
}

// Test LongestCommonPrefix solution
func TestLongestCommonPrefix(t *testing.T) {
	input1 := []string{"flower", "flow", "flight"}
	expected1 := "fl"
	actual1 := solutions.LongestCommonPrefix(input1)
	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Result was incorrect for LongestCommonPrefix test1, got: %s, exepected: %s.", actual1, expected1)
	}
	input2 := []string{"dog", "racecar", "car"}
	expected2 := ""
	actual2 := solutions.LongestCommonPrefix(input2)
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Result was incorrect for LongestCommonPrefix test2, got: %s, exepected: %s.", actual2, expected2)
	}
}

// Test GetConcatenation solution
func TestGetConcatenation(t *testing.T) {
	input1 := []int{1, 2, 1}
	expected1 := []int{1, 2, 1, 1, 2, 1}
	actual1 := solutions.GetConcatenation(input1)
	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Result was incorrect for GetConcatenation test1, got: %v, exepected: %v.", actual1, expected1)
	}
	input2 := []int{1, 3, 2, 1}
	expected2 := []int{1, 3, 2, 1, 1, 3, 2, 1}
	actual2 := solutions.GetConcatenation(input2)
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Result was incorrect for GetConcatenation test2, got: %v, exepected: %v.", actual2, expected2)
	}
}

// Test BuildArray solution
func TestBuildArray(t *testing.T) {
	input1 := []int{0, 2, 1, 5, 3, 4}
	expected1 := []int{0, 1, 2, 4, 5, 3}
	actual1 := solutions.BuildArray(input1)
	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Result was incorrect for BuildArray test1, got: %v, exepected: %v.", actual1, expected1)
	}
	input2 := []int{5, 0, 1, 2, 3, 4}
	expected2 := []int{4, 5, 0, 1, 2, 3}
	actual2 := solutions.BuildArray(input2)
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Result was incorrect for BuildArray test2, got: %v, exepected: %v.", actual2, expected2)
	}
}

// Test ConvertTemperature solution
func TestConvertTemperature(t *testing.T) {
	input1 := 36.50
	expected1 := []float64{309.65000, 97.70000}
	actual1 := solutions.ConvertTemperature(input1)
	if !reflect.DeepEqual(expected1, actual1) {
		t.Errorf("Result was incorrect for ConvertTemperature test1, got: %v, exepected: %v.", actual1, expected1)
	}
	input2 := 122.11
	expected2 := []float64{395.26000, 251.79800}
	actual2 := solutions.ConvertTemperature(input2)
	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Result was incorrect for ConvertTemperature test2, got: %v, exepected: %v.", actual2, expected2)
	}
}
