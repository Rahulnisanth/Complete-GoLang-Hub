package main

// COUNT THE NUMBER OF BITWISE OR SUBSETS :
func countMaxOrSubsets(nums []int) int {
	max_or := 0
	for _, num := range nums {
		max_or |= num
	}
	return backtrack(max_or, &nums, 0, 0)
}

func backtrack(max_or int, nums *[]int, idx int, curr int) int {
	if idx == len(*nums) {
		if max_or == curr {
			return 1
		}
		return 0
	} else {
		include := backtrack(max_or, nums, idx+1, curr|(*nums)[idx])
		exclude := backtrack(max_or, nums, idx+1, curr)
		return (include + exclude)
	}
}
