func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, tok := range tokens {
		if v, err := strconv.Atoi(tok); err == nil {
			stack = append(stack, v)
			continue
		}

		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch tok {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			stack = append(stack, a/b)
		}
	}

	return stack[0]
}
