// Задание 1. Работа с файлами
// Что нужно сделать
// Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение в формате:

// 2020-02-10 15:00:00 продам гараж.

// При вводе слова exit программа завершает работу.

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Не удалось создать или открыть файл", err)
		panic(err)
	}
	defer logFile.Close()

	var inputMessage string
	for {
		fmt.Print("Введите сообщение для сохранения в файл (exit - для выхода): ")
		inputMessage = readStringTerminal()
		if inputMessage == "exit" {
			return
		}
		logTime := time.Now()
		logFile.WriteString(fmt.Sprintf("%v %s\n", logTime.Format("2006-01-02 15:04:05"), inputMessage))
	}
}

func readStringTerminal() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return input.Text()
}
