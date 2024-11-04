package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// Task Base structure for tasks
type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// GetTasks A function to get task list
func GetTasks() ([]Task, int) {
	data, err := os.ReadFile("storage.json")
	if err != nil {
		panic(err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return []Task{}, 0
	}
	return tasks, 1

}

// UpdateTasks A function to update task list
func UpdateTasks(updatedTasks []Task, file *os.File) {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err := os.WriteFile("storage.json", []byte(""), 0644)
	if err != nil {
		panic(err)
	}
	for index, task := range updatedTasks {

		if index == 0 {
			_, err = file.WriteString("[\n")
			if err != nil {
				panic(err)
			}
		}
		err = encoder.Encode(task)
		if err != nil {
			panic(err)
		}
		if index == len(updatedTasks)-1 {
			_, err = file.WriteString("]\n")
			if err != nil {
				panic(err)
			}
		} else {
			_, err = file.WriteString(",")
			if err != nil {
				panic(err)
			}
		}
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}

}

// ClearScreen A function to clear the screen
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Используйте cmd для Windows
	} else {
		cmd = exec.Command("clear") // Команда "clear" для Unix
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}

//CreateTask A function to create new task

func CreateTask() {
	ClearScreen()
	var Id int
	namePrompt := promptui.Prompt{
		Label:     "Enter a task name",
		Templates: &promptui.PromptTemplates{},
		AllowEdit: true, // Позволяет редактировать введенный текст
	}
	taskName, err := namePrompt.Run()
	if err != nil {
		fmt.Printf("Failed: %v", err)

	}
	ClearScreen()
	descriptionPrompt := promptui.Prompt{
		Label:     "Enter a description of the task",
		Templates: &promptui.PromptTemplates{},
		AllowEdit: true,
	}
	taskDescription, err := descriptionPrompt.Run()
	if err != nil {
		fmt.Printf("Failed", err)

	}
	file, err := os.OpenFile("storage.json", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	tasks, status := GetTasks()
	//os.WriteFile("storage.json", []byte(""), 0644)
	if status == 1 {
		Id = tasks[len(tasks)-1].Id + 1
	} else {
		Id = 1
	}

	newTask := Task{Id, taskName, taskDescription, false}
	tasks = append(tasks, newTask)
	UpdateTasks(tasks, file)
	ClearScreen()
	ShowTasks()

}

func ShowTasks() {
	ClearScreen()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	tasks, status := GetTasks()
	if status == 1 {
		for _, task := range tasks {
			if task.Done {

				fmt.Println(" Task #", task.Id, "*********************************************", "\n", "Name \t \t ", task.Name, "\n", "Description\t ", task.Description, "\n", "Done?\t \t ", green("DONE"))
			} else {

				fmt.Println(" Task #", task.Id, "*********************************************", "\n", "Name \t \t ", task.Name, "\n", "Description\t ", task.Description, "\n", "Done?\t \t ", red("UNDONE"))
			}
		}
	} else {
		fmt.Println("No tasks yet")
	}

}

// DeleteTask A function to delete a task from the list
func DeleteTask() {
	file, err := os.OpenFile("storage.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	tasks, state := GetTasks()
	if state == 1 {
		ShowTasks()

		fmt.Println("Enter the id of the task you want to delete (0 to get back)")
		var toDelete int
		_, err = fmt.Scan(&toDelete)
		if err != nil {
			panic(err)
		}
		if toDelete == 0 {
			ClearScreen()
			ShowTasks()
			return
		}
		for index, task := range tasks {
			if task.Id == toDelete {
				tasks = append(tasks[:index], tasks[index+1:]...)
				break
			}
		}
		UpdateTasks(tasks, file)
		ClearScreen()
		ShowTasks()
	} else {
		ClearScreen()
		fmt.Println("No tasks yet")
	}

}

// ChangeDone A function to change status of the task
func ChangeDone() {

	tasks, state := GetTasks()
	if state == 1 {
		ShowTasks()
		fmt.Println("Enter the task id, the status you want to change")
		taskId := 0
		_, err := fmt.Scan(&taskId)
		if err != nil {
			panic(err)
		}

		fmt.Println("Enter the status:\n1 - Done\n2 - Undone")
		status := 0
		_, err = fmt.Scan(&status)
		if err != nil {
			panic(err)
		}
		for index, task := range tasks {
			if task.Id == taskId {
				if status == 1 {
					tasks[index].Done = true
				} else {
					tasks[index].Done = false
				}
			} else {
				continue
			}
		}
		file, err := os.OpenFile("storage.json", os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}
		UpdateTasks(tasks, file)
		ClearScreen()
		ShowTasks()

	} else {
		ClearScreen()
		fmt.Println("No tasks yet")
	}

}
