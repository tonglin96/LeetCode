package _go

func isValid(s string) bool {
	var flagMap = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if _, ok := flagMap[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != flagMap[s[i]] {
				return false
			}
			stack = stack[len(stack):]
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
