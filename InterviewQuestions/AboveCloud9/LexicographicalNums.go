package main

// Function to generate lexicographical numbers
func lexicalOrder(n int) []int {
	result := []int{}
	for i := 1; i <= 9; i++ {
		if i <= n {
			dfs(n, i, &result)
		}
	}
	return result
}

// Helper function to perform DFS
func dfs(n int, num int, result *[]int) {
	if num > n {
		return
	}
	*result = append(*result, num)
	for i := 0; i <= 9; i++ {
		next_num := (num * 10) + i
		if next_num > n {
			break
		}
		dfs(n, next_num, result)
	}
}
