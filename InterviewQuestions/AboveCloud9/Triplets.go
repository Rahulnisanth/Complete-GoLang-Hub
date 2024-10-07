package main

func findTriplets(arr ...int) [][]int {
	N := len(arr)
	result := [][]int{}
	for i := 0; i < N-2; i += 1 {
		for j := i + 1; j < N-1; j += 1 {
			for k := j + 1; k < N; k += 1 {
				if arr[i]+arr[j]+arr[k] == 0 {
					result = append(result, []int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}
	return result
}
