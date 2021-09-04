package cmd

import (
	"fmt"
	"os"
	"tasks/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: " List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := db.IncompleteTask()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(list) == 0 {
			fmt.Println("There are no Task in your list")
			return
		}
		fmt.Println("These are the tasks in your to do list")
		for i, task := range list {
			fmt.Printf("%d: %s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
