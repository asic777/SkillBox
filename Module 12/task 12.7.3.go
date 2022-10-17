// Задание 3. Уровни доступа
// Что нужно сделать
// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.

// Рекомендация
// Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

const fileName = "accessLevels.txt"

func main() {
	fmt.Println("-------------------")
	fmt.Println("Создаем новый файл:")
	CreateFileReadOnly(fileName, 0444)
	fmt.Println("-------------------")
	fmt.Println("Данные из файла:")
	fmt.Println(ReadFromFile(fileName))
	fmt.Println("-------------------")
	fmt.Println("Проверка записи в файл только для чтения:")
	WriteStringToFile(fileName, "Попытка записи в файл только для чтения\n")
	fmt.Println("-------------------")
}

func CreateFileReadOnly(nameFile string, mode fs.FileMode) {
	osFileHandler, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Не удалось создать или открыть файл", err)
		//panic(err)
		return
	}
	defer osFileHandler.Close()

	if os.Chmod(nameFile, mode) != nil {
		fmt.Println("Не удалось изменить права доступа к файлу", err)
		//panic(err)
	}

	if _, err := osFileHandler.WriteString("Файл только для чтения\n"); err != nil {
		fmt.Println("Не удалось записать в файл строку", err)
		//panic(err)
	}
	fmt.Printf("Файл %s успешно создан\n", fileName)
}
func ReadFromFile(nameFile string) string {
	osFileHandler, err := os.Open(nameFile)
	if err != nil {
		fmt.Println("Не удалось открыть файл для чтения данных", err)
		//panic(err)
		return err.Error()
	}
	defer osFileHandler.Close()

	infoFile, err := os.Stat(nameFile)
	if err != nil {
		fmt.Println("Не удалось прочитать информацию о файле", err)
		//panic(err)
	}

	bufferForRead := make([]byte, infoFile.Size())
	if _, err := io.ReadFull(osFileHandler, bufferForRead); err != nil {
		fmt.Println("Не удалось прочитать последовательность байтов из файла", err)
		//panic(err)
	}
	return string(bufferForRead)
}

func WriteStringToFile(nameFile string, text string) {
	osFileHandler, err := os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		fmt.Println("Не удалось открыть файл для записи", err)
		//panic(err)
		return
	}
	defer osFileHandler.Close()

	if _, err := osFileHandler.WriteString(text); err != nil {
		fmt.Println("Не удалось записать в файл строку", err)
		//panic(err)
	}
}
