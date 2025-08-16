/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/AaronBrownDev/go-todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		title, err := cmd.Flags().GetString("title")
		if err != nil {
			log.Fatalln(err)
		}
		description, err := cmd.Flags().GetString("desc")
		if err != nil {
			log.Fatalln(err)
		}
		priority, err := cmd.Flags().GetString("priority")
		if err != nil {
			log.Fatalln(err)
		}
		category, err := cmd.Flags().GetString("category")
		if err != nil {
			log.Fatalln(err)
		}
		due, err := cmd.Flags().GetTime("due")
		if err != nil {
			log.Fatalln(err)
		}

		jsonRepo.Save(&todo.Todo{
			ID:          "0",
			Title:       title,
			Description: description,
			Priority:    todo.Priority(priority),
			Category:    category,
			DueDate:     due,
			Completed:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Time{},
		})

	},
}

func init() {
	addCmd.Flags().String("title", "", "")
	addCmd.Flags().String("desc", "", "")
	addCmd.Flags().String("priority", "", "")
	addCmd.Flags().String("category", "", "")
	addCmd.Flags().Time("due", time.Time{}, []string{}, "")

	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
