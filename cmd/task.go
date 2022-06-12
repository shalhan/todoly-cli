/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"todoly.app/todoly/ui"
	"todoly.app/todoly/task/model"
	"todoly.app/todoly/task/repo/sqlite"
	"github.com/spf13/cobra"
)

var taskRepo *sqlite.SQLiteRepository = sqlite.NewTaskRepo()
var taskDeleteId string
var taskTitle string
var taskBrief string


// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Show, create and delete task",
	Run: func(cmd *cobra.Command, args []string) {
		activeBoard, err := boardRepo.GetActive()
		if err != nil {
			log.Fatal(err)
			return
		}
		if taskDeleteId != "" && taskTitle != "" {
			fmt.Println("Err: you can`t create and delete task at the same time")
			return
		}
		if taskDeleteId != "" {
			err := taskRepo.Delete(append(args, taskDeleteId))
			if (err != nil) {
				log.Fatal(err)
			}
		} else if taskTitle != "" {
			task := model.Task{Title: taskTitle, Brief: taskBrief, BoardId: activeBoard.ID} 
			_, err := taskRepo.Create(task) 
			if err != nil {
				log.Fatal(err)
			}
		}

		all, err := taskRepo.All(activeBoard.ID)
		// print result
		ui.PrintTaskList(all)
		return;
	},
}


func init() {
	taskCmd.PersistentFlags().StringVarP(&taskTitle, "new", "n", "", "Create new task")
	taskCmd.PersistentFlags().StringVarP(&taskBrief, "brief", "b", "", "Add brief of the task")

	taskCmd.PersistentFlags().StringVarP(&taskDeleteId, "delete", "d", "", "Delete task")

	rootCmd.AddCommand(taskCmd)
}
