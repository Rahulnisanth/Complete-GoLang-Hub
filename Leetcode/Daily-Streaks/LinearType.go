package main

// MINIMUM NUMBER OF SWAPS TO MAKE THE STRING BALANCED :
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


// MINIMUM ADD TO MAKE PARENTHESIS BALANCED :
func minAddToMakeValid(s string) int {
    open := 0
    close_needed := 0
    for _, ch := range s {
        if ch == '(' {
            open++
        } else if ch == ')' {
            if open > 0 {
                open--
            } else {
                close_needed++
            }
        }
    }
    return open + close_needed
}


// SEPERATE THE WHITE AND BLACK BALLS :
func minimumSteps(s string) int64 {
    var result, zeroes int64
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			zeroes++
		} else if s[i] == '1' {
			result += zeroes
		}
	}
	return result
}


// MAXIMUM SWAPS TO MAKE THE INTEGER LARGEST :
func maximumSwap(num int) int {
	nums := []rune(strconv.Itoa(num))
	for i := 0; i < len(nums); i++ {
		idx := i
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[idx] {
				idx = j
			}
		}
		if idx != i && nums[i] < nums[idx] {
			nums[i], nums[idx] = nums[idx], nums[i]
			result, _ := strconv.Atoi(string(nums))
			return result
		}
	}
	return num
}