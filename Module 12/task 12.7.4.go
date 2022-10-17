// Задание 4. Пакет ioutil
// Что нужно сделать
// Перепишите задачи 1 и 2, используя пакет ioutil.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const fileName = "logioutil.txt"

func main() {
	var bufferWriteToFile bytes.Buffer
	var inputMessage string
	for {
		fmt.Print("Введите сообщение для сохранения в файл (exit - для выхода): ")
		inputMessage = readStringTerminal()
		if inputMessage == "exit" {
			break
		}
		logTime := time.Now()
		bufferWriteToFile.WriteString(fmt.Sprintf("%v %s\n", logTime.Format("2006-01-02 15:04:05"), inputMessage))
	}
	err := ioutil.WriteFile(fileName, bufferWriteToFile.Bytes(), 0777)
	if err != nil {
		fmt.Println("Не удалось записать в файл последовательность байтов", err)
		//panic(err)
	}
	fmt.Println(ReadFromFile(fileName))
}

func readStringTerminal() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	if input.Err() != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", input.Err())
	}
	return input.Text()
}

func ReadFromFile(nameFile string) string {
	osFileHandler, err := os.Open(nameFile)
	if err != nil {
		fmt.Println("Не удалось открыть файл для чтения данных")
		//panic(err)
		return err.Error()
	}
	defer osFileHandler.Close()

	bufferForRead, err := ioutil.ReadAll(osFileHandler)
	if err != nil {
		fmt.Println("Не удалось прочитать последовательность байтов из файла", err)
		//panic(err)
	}
	return string(bufferForRead)
}
