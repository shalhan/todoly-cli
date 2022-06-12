/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"todoly.app/todoly/ui"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Command to set task as done",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		activeBoard, err := boardRepo.GetActive()
		if err != nil {
			log.Fatal(err)
			return
		}

		err2 := taskRepo.SetDone(args)
		if err2 != nil {
			log.Fatal(err2)
		}
		all, err := taskRepo.All(activeBoard.ID)
		ui.PrintTaskList(all)
		return;

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
