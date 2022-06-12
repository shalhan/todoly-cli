package ui

import (
	"fmt"

	"todoly.app/todoly/task/model"
)

func PrintTaskList(tasks []model.Task) {
	for _, task := range tasks {
		taskTitle := task.Title
		taskStatus := " "
		if task.IsActive {
			taskTitle = greenFont(taskTitle)
		}
		if task.Status == "Done" {
			taskStatus = greenFont("o") 
		} else if task.Status == "Abort" {
			taskStatus = redFont("x")
		} else if task.Status == "Going" {
			taskStatus = yellowFont("~")
		}

		fmt.Print("\n")
		fmt.Printf(" [%s]  %d  %s \n", taskStatus, task.ID, taskTitle)
		if task.Brief != "" {
			fmt.Printf("  %s %s\n", cyanFont("|_"), lightGrayFont(task.Brief))
		}
	}

}

func greenFont(str string) string {
	return "\033[0;32m" + str + "\033[39m"
}

func redFont(str string) string {
	return "\033[0;31m" + str + "\033[0m"
}

func yellowFont(str string) string {
	return "\033[0;33m" + str + "\033[0m"
}

func lightGrayFont(str string) string {
	return "\033[0;37m" + str + "\033[39m"
}

func cyanFont(str string) string {
	return "\033[0;36m" + str + "\033[0m"
}


