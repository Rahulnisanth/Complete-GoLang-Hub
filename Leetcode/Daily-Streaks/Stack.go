// MIN STRING LENGTH AFTER THE REMOVAL OF SUB-STRINGS :
package dailystreaks

func minLength(s string) int {
    stack := []rune{}
    for _, ch := range s {
        if len(stack) > 0 {
            if stack[len(stack) - 1] == 'A' && ch == 'B' {
                stack = stack[:len(stack) - 1]
            } else if stack[len(stack) - 1] == 'C' && ch == 'D' {
                stack = stack[:len(stack) - 1]
            } else {
                stack = append(stack, ch)
            }
        } else {
            stack = append(stack, ch)
        }
    }
    return len(stack)
}