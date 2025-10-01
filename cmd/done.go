package cmd

import (
	"fmt"
	"strconv"

	"github.com/snowchest67/todo.git/todo"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [номер задачи]",
	Short: "Выполнить задачу",
	Long:  `Пометить выбранную задачу как выполненную`,
	Args:  cobra.ExactArgs(1), 
	RunE: func(cmd *cobra.Command, args []string) error {
		taskDone := args[0]

		tasks, err := todo.LoadJSONFile("tasklist.json")
		if err != nil{
			return err
		}

		err = tasks.CompleteTask(taskDone)
		if err != nil{
			return err
		}

		err = todo.SaveTaskList("tasklist.json", tasks)
		if err != nil{
			return err
		}

		num, err := strconv.Atoi(taskDone)
		if err != nil {
		return err
	}
		fmt.Printf("✅ Задача выполнена: %s\n", (*tasks)[num-1].Title)
		return nil
	},
}

func init(){
	rootCmd.AddCommand(doneCmd)
}