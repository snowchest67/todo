package todo

import "testing"

func TestTaskList_AppendTask(t *testing.T){
tests := []struct{
	name string
	a []Task
	expected int
}{
	{"Добавление одной задачи", []Task{{Title: "Купить молоко", Done: false}}, 2},
	{"Добавление нескольких задач", []Task{{Title: "Купить молоко", Done: false}, {Title: "Убрать за котом", Done: true}}, 3},
	{"Добавление пустого списка задач", []Task{}, 1},
}

for _, tt := range tests{
	t.Run(tt.name, func(t *testing.T){
		tasks := TaskList{{Title: "Купить молоко", Done: false}}
		tasks.AppendTask(tt.a...)
		result := len(tasks)
		if tt.expected != result {
		t.Errorf("len(tasks) = %d; expected %d", result, tt.expected)
	}
})
}
}

func TestTaskList_CompleteTask(t *testing.T){
	tests := []struct{
		name string
		a string
		expected bool
	}{
		{"Корректное завершение задачи", "1", true},
		{"Неверный индекс (меньше 1)", "-1", false},
		{"Неверный индекс (больше длины)", "3", false},
		{"Некорректный ввод (не число)", "dfe", false},
	}

	for _, tt := range tests{
	t.Run(tt.name, func(t *testing.T){
		tasks := TaskList{{Title: "Купить молоко", Done: false}}
		err := tasks.CompleteTask(tt.a)
		if err != nil{
			t.Error(err)
		}
		result := tasks[0].Done
		if tt.expected != result && err == nil  {
		t.Errorf("len(tasks) = %t; expected %t", result, tt.expected)
	}
})
}


}