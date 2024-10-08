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