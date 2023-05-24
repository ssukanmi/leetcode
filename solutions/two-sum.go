package solutions

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

func TwoSum(nums []int, target int) []int {
	output := make([]int, 2)
out:
	for idx1, val1 := range nums {
		for idx2, val2 := range nums {
			if idx1 != idx2 && val1+val2 == target {
				output = []int{idx1, idx2}
				break out
			}
		}
	}
	return output
}

func TwoSum2(nums []int, target int) []int {
	output := make([]int, 2)
	foundVals := make(map[int]int)

	for idx, val := range nums {
		compliment := target - val
		// checks if compliment is in found values
		if v, ok := foundVals[compliment]; ok {
			output = []int{v, idx}
			break
		}
		foundVals[val] = idx
	}
	return output
}
