/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cbrissonCoveo/task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DelTask(task.ID)
			if err != nil {
				fmt.Printf("Failed to make \"%d\" as complete. Error: %s\n", id, err)
			} else {
				fmt.Printf("Task \"%d\" was deleted succesfully.\n", id)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)

}
