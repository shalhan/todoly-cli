/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"todoly.app/todoly/ui"
	"github.com/spf13/cobra"
)

// goingCmd represents the going command
var goingCmd = &cobra.Command{
	Use:   "going",
	Short: "Command to set task as in progress",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		activeBoard, err := boardRepo.GetActive()
		if err != nil {
			log.Fatal(err)
			return
		}

		err2 := taskRepo.SetGoing(args)
		if err2 != nil {
			log.Fatal(err2)
		}
		all, err := taskRepo.All(activeBoard.ID)
		ui.PrintTaskList(all)	
		return;

	},
}

func init() {
	rootCmd.AddCommand(goingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
