package cmd

import (
	"fmt"
	"time"

	"github.com/snowchest67/todo.git/todo"

	"github.com/spf13/cobra"
)

var (
	priority int
	dueDate  string
)

var addCmd = &cobra.Command{
	Use:   "add [текст задачи]",
	Short: "Добавить новую задачу",
	Long:  `Добавить новую задачу в список дел с возможностью указать приоритет и дату выполнения.`,
	Args:  cobra.ExactArgs(1), 
	RunE: func(cmd *cobra.Command, args []string) error {
		taskText := args[0]

		if priority < 1 || priority > 5{
			return fmt.Errorf("приоритет должен быть от 1 до 5")
		}

		if dueDate != ""{
			if _ , err := time.Parse("2006-01-02", dueDate); err != nil{
				return fmt.Errorf("неверный формат даты: используйте YYYY-MM-DD")
			}
		}

		tasks, err := todo.LoadJSONFile("tasklist.json")
		if err != nil{
			return err
		}

		tasks.AppendTask(todo.Task{Title: taskText, Done: false, Priority: priority, DueDate: dueDate})

		err = todo.SaveTaskList("tasklist.json", tasks)
		if err != nil{
			return err
		}

		fmt.Printf("Задача добавлена: %s\n", taskText)
		return nil
	},
}

func init(){
	addCmd.Flags().IntVarP(&priority, "priority", "p", 1, "Приоритет (1-5)")
	addCmd.Flags().StringVarP(&dueDate, "due", "d", "", "Дата выполнения (YYYY-MM-DD)")
	rootCmd.AddCommand(addCmd)
}