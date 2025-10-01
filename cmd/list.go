package cmd

import (
	"fmt"
	"sort"

	"github.com/snowchest67/todo.git/todo"

	"github.com/spf13/cobra"
)

var flagSort string

var listCmd = &cobra.Command{
	Use:   "list [--sort=priority|--sort=due]",
	Short: "Показать все задачи",
	Long:  "Выводит список всех задач. Можно сортировать с помощью --sort=priority или --sort=due.",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := todo.LoadJSONFile("tasklist.json")
		if err != nil{
			return err
		}

		if len(*tasks) == 0 {
			tasks.PrintTasks()
			return nil
		}

		switch flagSort{
			case "priority":
			sort.Slice(*tasks, func(i, j int) bool {
				return (*tasks)[i].Priority > (*tasks)[j].Priority
			})
			case "due":
				sort.Slice(*tasks, func(i,j int) bool{
					if (*tasks)[i].DueDate == "" && (*tasks)[j].DueDate == ""{
						return false
					}
					if (*tasks)[i].DueDate == "" {
						return false
					}
					if (*tasks)[j].DueDate == "" {
						return true
					}
					return (*tasks)[i].DueDate < (*tasks)[j].DueDate
				})
			case "":

			default:
				return fmt.Errorf("неизвестный тип сортировки: %s (допустимо: priority, due)", flagSort)
			}

		tasks.PrintTasks()
		return nil
},
}

func init(){
	listCmd.Flags().StringVarP(&flagSort, "sort", "s", "", "Сортировка")
	rootCmd.AddCommand(listCmd)
}