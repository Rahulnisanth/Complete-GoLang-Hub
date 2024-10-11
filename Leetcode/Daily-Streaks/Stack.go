// MIN STRING LENGTH AFTER THE REMOVAL OF SUB-STRINGS :
package main

func minLength(s string) int {
	stack := []rune{}
	for _, ch := range s {
		if len(stack) > 0 {
			if stack[len(stack)-1] == 'A' && ch == 'B' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] == 'C' && ch == 'D' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, ch)
			}
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

// MAXIMUM WIDTH RAMP :
func maxWidthRamp(nums []int) int {
	N := len(nums)
	stack := make([]int, N)
	for i := 0; i < N; i++ {
		if len(stack) == 0 || nums[i] < nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	result := 0
	for j := N - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[j] >= nums[stack[len(stack)-1]] {
			result = max(result, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			if result == j {
				return result
			}
		}
	}
	return result
}
