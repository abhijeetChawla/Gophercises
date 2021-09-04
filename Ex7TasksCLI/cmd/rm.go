package cmd

import (
	"fmt"
	"os"
	"strconv"
	"tasks/db"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Will delete Task from the List",
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

		t, err := db.DeleteTask(list[taskIndex-1].Key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("You have deleted the '%s' task.", t.Value)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
