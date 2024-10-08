package main

// 
func minSwaps(s string) int {
    count := 0
    for _, ch := range s {
        if ch == '[' {
            count += 1
        } else if count > 0 {
            count -= 1
        }
    }
    return (count + 1) / 2
}