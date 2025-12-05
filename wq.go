package main

import (
	"fmt"
)

type Task struct {
	Oglavlenie string
	Vipolnenie bool
}

func main() {
	var tasks []Task
	var choice int

	for {

		fmt.Println("---to-do list---")

		fmt.Println("1.Добавить действие")

		fmt.Println("2.Посмотреть действия")

		fmt.Println("3.Ометить выполненой")

		fmt.Println("4.Выход")

		fmt.Print("Выебрите действие: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:

			var oglavlenie string
			fmt.Println("Введите название вашего действия:")
			fmt.Scan(&oglavlenie)
			tasks = append(tasks, Task{Oglavlenie: oglavlenie})
			fmt.Println("Действие добавлено успешно.")

		case 2:

			fmt.Println("\nСписок действий: ")
			if len(tasks) == 0 {
				fmt.Println("\nПусто")
			}
			for i, t := range tasks {
				status := "..."
				if t.Vipolnenie {
					status = "good"
				}
				fmt.Println(i+1, status, t.Oglavlenie)
			}

		case 3:
			var num int
			fmt.Println("Номер задачи: ")
			fmt.Scan(&num)

			if num < 1 || num > len(tasks) {
				fmt.Println("Неверный номер действия")
				continue

			}

			tasks[num-1].Vipolnenie = true
			fmt.Println("good, отмечено как выполненое действие.")

		case 4:

			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте еще")
		}
	}

}
