// Задача №1
// Цель практической работы
// Закрепить навык работы в редакторе JetBrains GoLand.
// Что нужно сделать
// В одном из модулей мы учились выводить шахматное поле в консоль, используя циклы.
// Произведите рефакторинг одного из вариантов решения этого задания. Выведите в функцию процесс вывода шахматного поля в консоль.
// Постарайтесь подобрать говорящие имена для функции и переменных вместо a, b, k и j.

// func main() {
//     fmt.Println("Введите ширину и высоту поля")
//     var a, b int
//     fmt.Scan(&a, &b)

//     for k := 0; k < b; k++ {
//         for j := 0; j < a; j++ {
//             if (k+j)%2 == 0 {
//                 fmt.Print("*")
//             } else {
//                 fmt.Print(" ")
//             }
//         }
//         fmt.Println()
//     }
//     fmt.Println("end")
// }

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите ширину и высоту поля:")
	var widthChessTable, heigthChessTable int
	fmt.Scan(&widthChessTable, &heigthChessTable)

	generateChessTable(widthChessTable, heigthChessTable)

	fmt.Println("end")
}

func generateChessTable(width, heigth int) {
	for rows := 0; rows < heigth; rows++ {
		for column := 0; column < width; column++ {
			if (column+rows)%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
