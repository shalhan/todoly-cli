/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"todoly.app/todoly/board/model"
	"todoly.app/todoly/board/repo/sqlite"
	"github.com/spf13/cobra"
)

var boardRepo *sqlite.SQLiteRepository = sqlite.NewBoardRepo()
var deleteId string
var newBoard string
// boardCmd represents the board command
var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "Show, create and delete board",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if  deleteId != "" && newBoard != "" {
			fmt.Println("Err: you can`t create and delete board at the same time")
			return
		}

		if deleteId != "" {
			err := boardRepo.Delete(deleteId)
			if (err != nil) {
				log.Fatal(err)
			}
		} else if newBoard != "" {
			name := newBoard

			board := model.Board{Name: name}

			if _, err := boardRepo.Create(board); err != nil {
				log.Fatal(err)
			}
		} else if len(args) > 0 {
			err := boardRepo.Active(args[0])
			if (err != nil) {
				log.Fatal(err)
			}
		}


		all, err := boardRepo.All()
		if err != nil {
			log.Fatal(err)
		}
		for _, board := range all {
			activeIcon := " "
			boardName := board.Name
			if board.IsActive {
				activeIcon = greenFont("*")
				boardName = greenFont(boardName)
			}
			fmt.Printf(" %s  %s\n", activeIcon, boardName)
		}
		return;


	},
}

func greenFont(str string) string {
	return "\033[0;32m" + str + "\033[39m"
}


func redFont(str string) string {
	return "\033[0;31m" + str + "\033[0m"
}

func init() {
	boardCmd.PersistentFlags().StringVarP(&newBoard, "new", "n", "", "Create new board")

	boardCmd.PersistentFlags().StringVarP(&deleteId, "delete", "d", "", "Delete existing board")

	rootCmd.AddCommand(boardCmd)
}
