package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"main/tasks"
	"os"
)

func main() {
	tasks.ClearScreen()
	fmt.Println("Hello there, it is hhellbentt`s simple task tracker to practice go. If u found bugs or errors plz contact me")
	items := []string{"show tasks", "create new task", "mark task as done(undone)", "delete task from lists", "exit"}
	for {
		prompt := promptui.Select{
			Label: "What to do?",
			Items: items,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You chose %q\n", result)
		switch result {
		case "show tasks":

			tasks.ShowTasks()
		case "create new task":

			tasks.CreateTask()
		case "mark task as done(undone)":

			tasks.ChangeDone()
		case "delete task from lists":

			tasks.DeleteTask()
		case "exit":

			os.Exit(1)
		}

	}
}
