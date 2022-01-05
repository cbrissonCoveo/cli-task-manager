/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cbrissonCoveo/task/db"
	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

type Task struct {
	ID   int
	name string
	done bool
}
type Database struct {
	db *bbolt.DB
}

func (d *Database) CreateTask(t *Task) error {
	return d.db.Update(func(tx *bbolt.Tx) error {

		return nil
	})
}

func init() {
	rootCmd.AddCommand(addCmd)

}
