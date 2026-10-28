package main

import "os"

func stringToInt(str string) (int, bool) {
	digit := int64(0)
	multiplier := int64(1)
	const maxInt64 = 9223372036854775807
	const minInt64 = -9223372036854775808

	for idx, char := range str {
		if idx == 0 && char == '-' {
			multiplier = -1
			continue
		}
		if char >= '0' && char <= '9' {
			val := int64(char - '0')
			if digit > (maxInt64-val)/10 {
				return 0, false
			}
			digit = (digit * 10) + val
		} else {
			return 0, false
		}
	}
	result := multiplier * digit
	if result < minInt64 || result > maxInt64 {
		return 0, false
	}

	return int(result), true
}

func isNumeric(str string) bool {
	for idx, char := range str {
		if idx == 0 && char == '-' {
			continue
		}
		if !(char >= '0' && char <= '9') {
			return false
		}
	}
	return true
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	result := intToStringHelper(n, []rune{}, isNegative)
	return result
}

func intToStringHelper(n int, arr []rune, isNegative bool) string {
	if n == 0 {
		str := ""
		if isNegative {
			str += "-"
		}
		for i := len(arr) - 1; i >= 0; i-- {
			str += string(arr[i])
		}
		return str
	}

	mod := n % 10
	rem := n / 10
	arr = append(arr, rune(mod+'0'))
	return intToStringHelper(rem, arr, isNegative)
}

func addOverflows(a, b int) bool {
	if b > 0 && a > 9223372036854775807-b {
		return true
	}
	if b < 0 && a < -9223372036854775808-b {
		return true
	}
	return false
}

func subOverflows(a, b int) bool {
	if b < 0 && a > 9223372036854775807+b {
		return true
	}
	if b > 0 && a < -9223372036854775808+b {
		return true
	}
	return false
}

func mulOverflows(a, b int) bool {
	if a == 0 || b == 0 {
		return false
	}
	ans := a * b
	if ans/a != b {
		return true
	}
	return false
}

func main() {
	userArguments := os.Args[1:]

	if len(userArguments) > 3 || len(userArguments) < 3 {
		return
	}

	if !isNumeric(userArguments[0]) || !isNumeric(userArguments[2]) {
		return
	}

	operandA, isvalidA := stringToInt(userArguments[0])
	operator := userArguments[1]
	operandB, isValidB := stringToInt(userArguments[2])

	if !isvalidA || !isValidB {
		return
	}

	switch operator {
	case "*":
		if mulOverflows(operandA, operandB) {
			return
		}
		os.Stdout.Write([]byte(intToString(operandA * operandB)))
		os.Stdout.Write([]byte("\n"))
	case "/":
		if operandB == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		if operandA == -9223372036854775808 && operandB == -1 {
			return
		}
		os.Stdout.Write([]byte(intToString(operandA / operandB)))
		os.Stdout.Write([]byte("\n"))
	case "-":
		// if subOverflows(operandA, operandB) {
		// 	return
		// }
		os.Stdout.Write([]byte(intToString(operandA - operandB)))
		os.Stdout.Write([]byte("\n"))
	case "+":
		if addOverflows(operandA, operandB) {
			return
		}
		os.Stdout.Write([]byte(intToString(operandA + operandB)))
		os.Stdout.Write([]byte("\n"))
	case "%":
		if operandB == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		if operandA == -9223372036854775808 && operandB == -1 {
			return
		}
		os.Stdout.Write([]byte(intToString(operandA % operandB)))
		os.Stdout.Write([]byte("\n"))
	default:
		return
	}
}
