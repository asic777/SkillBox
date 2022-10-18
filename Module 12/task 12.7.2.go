// Задание 2. Интерфейс io.Reader
// Что нужно сделать
// Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.

// Рекомендация
// Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.

package main

import (
	"fmt"
	"io"
	"os"
)

const fileName = "log.txt"

func main() {
	logFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Не удалось открыть файл для чтения данных", err)
		//panic(err)
		return
	}
	defer logFile.Close()

	infoFile, err := logFile.Stat()
	if err != nil {
		fmt.Println("Не удалось прочитать информацию о файле", err)
		//panic(err)
	}

	bufferForRead := make([]byte, infoFile.Size())
	if _, err := io.ReadFull(logFile, bufferForRead); err != nil {
		fmt.Println("Не удалось прочитать последовательность байтов из файла", err)
		//panic(err)
	}
	fmt.Printf("Вывод данных из файла %s:\n", fileName)
	fmt.Println(string(bufferForRead))
}
