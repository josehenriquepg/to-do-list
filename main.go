package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks []Task
var nextID int

func main() {
	loadTasks()
	defer saveTasks()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMenu\n1. Add Task\n2. List Tasks\n3. Complete Task\n4. Remove Task\n5. Exit")

		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTask(scanner)
		case "2":
			listTasks()
		case "3":
			completeTask(scanner)
		case "4":
			removeTask(scanner)
		case "5":
			fmt.Println("Leaving...")
			return
		default:
			fmt.Println("Invalid option, choose again.")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Task description: ")
	scanner.Scan()
	description := scanner.Text()

	task := Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
	}
	nextID++

	tasks = append(tasks, task)
	fmt.Println("Task added!")
}

func listTasks() {
	for _, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, status)
	}
}

func completeTask(scanner *bufio.Scanner) {
	fmt.Print("Complete Task\nEnter the ID of the completed task: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task completed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func removeTask(scanner *bufio.Scanner) {
	fmt.Print("Remove Task\nEnter the task ID to remove it: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task removed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func saveTasks() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Erro ao salvar tarefas: ", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error when saving tasks:", err)
	}
}

func loadTasks() {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error loading tasks:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error when decoding tasks:", err)
		return
	}

	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
}
