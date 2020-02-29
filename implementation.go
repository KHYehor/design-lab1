package lab1

import (
	"errors"
	"fmt"
	"strings"
)

func isOperator(c byte) bool {
	return strings.ContainsAny(string(c), "+&-&*&/&^")
}

func isOperand(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isValid(str string) (bool, string) {
	countOperands := 0
	countOperators := 0
	length := len(str)
	fmt.Println(str == "")
	if str == "" {
		return false, "must be non empty string"
	}
	for i := 0; i < length; i++ {

		if i == 0 {
			if isOperand(str[i]) {
				if !isOperand(str[i+1]) {
					countOperands++
				}
				continue
			} else if isOperator(str[i]) {
				return false, "first char must be operand not operator"
			} else {
				return false, "first char must be operand not whitespace"
			}
		}

		if i == length-1 {
			if isOperator(str[i]) {
				countOperators++
				break
			} else if isOperand(str[i]) {
				countOperands++
				break
			} else {
				return false, "last char must be operand or operator"
			}
		}

		if string(str[i]) == " " {
			if isOperand(str[i+1]) || isOperator(str[i+1]) {
				continue
			} else {
				return false, "after whitespace must be operand or operator"
			}
		}

		if isOperand(str[i]) {
			if string(str[i+1]) == " " {
				countOperands++
				continue
			} else if isOperand(str[i+1]) {
				continue
			} else {
				return false, "after operand must be whitespace"
			}
		}

		if isOperator(str[i]) {
			countOperators++
			if string(str[i+1]) == " " {
				continue
			} else {
				return false, "after operator must be whitespace"
			}
		}
	}

	if countOperands != countOperators+1 {
		return false, "count of operators must be less operands then one"
	}
	return true, ""

}
func PostfixToPrefix(s string) (string, error) {

	valid, err := isValid(s)
	if !valid {
		fmt.Println(err)
		return "", errors.New(err)
	}
	var stack Stack
	length := len(s)
	temp := ""
	for i := 0; i < length; i++ {
		char := string(s[i])

		if char == " " {
			continue
		}

		if isOperator(s[i]) {
			op1 := stack.Top().(string)
			stack.Pop()
			op2 := stack.Top().(string)
			stack.Pop()

			stack.Push(char + " " + op2 + " " + op1)

		} else {
			temp += char
			if isOperand(s[i+1]) {
				continue
			} else {
				stack.Push(temp)
				temp = ""
			}

		}
	}
	return stack.Top().(string), nil
}
