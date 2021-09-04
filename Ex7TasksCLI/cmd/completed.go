package cmd

import (
	"fmt"
	"tasks/db"
	"time"

	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		maxTime := time.Now().Format(time.RFC3339)
		minTime := time.Now().AddDate(0, 0, -1).Format(time.RFC3339)
		fmt.Println(minTime, maxTime)
		list, err := db.CompletedTasks(minTime, maxTime)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(list) == 0 {
			fmt.Println("There are no completed tasks")
			return
		}
		fmt.Println("These are the tasks that have been compelted")
		for i, task := range list {
			fmt.Printf("%d: %s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
