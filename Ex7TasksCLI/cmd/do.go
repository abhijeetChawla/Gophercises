package cmd

import (
	"fmt"
	"os"
	"strconv"
	"tasks/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("This command requires and argument to work")
			return
		}

		taskIndex, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("The argument passed is not a number")
			return
		}

		list, err := db.IncompleteTask()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if taskIndex == 0 || taskIndex >= len(list) {
			fmt.Println("There is no task at that index. Try the list command to get the index")
			os.Exit(1)
		}
		t, err := db.DoTask(list[taskIndex-1].Key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("You have completed the '%s' task.", t.Value)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
