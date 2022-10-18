// Задание 5 (по желанию). Комбинации круглых скобок
// Что нужно сделать
// Напишите программу, которая на вход принимала бы интовое число и для него генерировала бы все возможные комбинации круглых скобок.

// Рекомендация
// Первый пример вывода программы:
// Введите количество пар скобок:
// 3
// ["((()))","(()())","(())()","()(())","()()()"]

// Второй пример вывода программы:
// Введите количество пар скобок:
// 1
// ["()"]

package main

import (
	"fmt"
)

func main() {
	fmt.Print("Введите количество пар скобок: ")
	var pairParentheses int
	fmt.Scan(&pairParentheses)

	fillParentheses(pairParentheses, 0, 0, "")
}
func fillParentheses(countPairParentheses int, openParentheses int, closeParentheses int, result string) {
	if closeParentheses == countPairParentheses {
		fmt.Printf("[%s]\n", result)
		return
	}
	if openParentheses < countPairParentheses {
		fillParentheses(countPairParentheses, openParentheses+1, closeParentheses, result+"(")
	}
	if openParentheses > closeParentheses {
		fillParentheses(countPairParentheses, openParentheses, closeParentheses+1, result+")")
	}
}
