package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Priority int `json:"priority"`
	DueDate string `json:"dueDate"`
}

type TaskList []Task

func (t *Task) priorityLabel() string {
	switch t.Priority {
	case 5:
		return "Критический"
	case 4:
		return "Высокий"
	case 3:
		return "Средний"
	case 2:
		return "Низкий"
	case 1:
		return "Тривиальный"
	default:
		if t.Priority < 1 {
			return "Критический"
		}
		return "Тривиальный"
	}
}

func (tl *TaskList) PrintTasks() {
	if len(*tl) == 0 {
		fmt.Println("Нет задач.")
		return
	}

	fmt.Println("Текущие задачи:")
	fmt.Println(strings.Repeat("-", 55))

	for i, task := range *tl {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}

		var parts []string
		parts = append(parts, fmt.Sprintf("%s %s", status, task.Title))

		if !task.Done {
			parts = append(parts, task.priorityLabel())
		}

		if task.DueDate != "" {
			parts = append(parts, fmt.Sprintf("Due: %s", task.DueDate))
		}

		line := strings.Join(parts, " ")
		fmt.Printf("%2d. %s\n", i+1, line)
	}

	fmt.Println(strings.Repeat("-", 55))

}

func (t *TaskList) AppendTask(newTasks ...Task) {
	*t = append(*t, newTasks...)
}

func (t *TaskList) CompleteTask(arg string) error {
	num, err := strconv.Atoi(arg)
	if err != nil {
		return err
	}
	index := num - 1
	if index < 0 || index >= len(*t) {
		return errors.New("неверный номер задачи")
	}
	(*t)[num-1].Done = true
	return nil
}

func LoadJSONFile(filename string) (*TaskList, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			defaultT := &TaskList{}
			if saveErr := SaveTaskList(filename, defaultT); saveErr != nil {
				return nil, fmt.Errorf("не удалось создать файл %s: %w", filename, saveErr)
			}
			return defaultT, nil
		}

		return nil, fmt.Errorf("не удалось открыть файл %s: %w", filename, err)
	}
	defer file.Close()

	var tasks TaskList
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("ошибка чтения JSON из %s: %w", filename, err)
	}

	return &tasks, nil

}

func SaveTaskList(filename string, tasks *TaskList) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}