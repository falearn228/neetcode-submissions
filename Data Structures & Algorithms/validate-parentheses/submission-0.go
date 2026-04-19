func isValid(s string) bool {
    if len(s) == 0 {   
        return true
    } 

    stack := make([]byte, 0, len(s))

    for i := range s {
        if len(stack) > 0 && match(stack[len(stack)-1], s[i]) {
            stack = stack[:len(stack)-1]
            continue
        }
        if isOpen(s[i]) {
            stack = append(stack, s[i])
        } else {
            return false
        }
    }

    return len(stack) == 0
}

func match(s byte, d byte) bool {
    if (s == '(' && d == ')') ||
    (s == '[' && d == ']') ||
    (s == '{' && d == '}') {
        return true
    }
    return false
}

func isOpen(s byte) bool {
    return s == '(' || s == '[' || s == '{'
}