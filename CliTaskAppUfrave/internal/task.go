package internal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	TODO_STATUS       = "todo"
	INPROGRESS_STATUS = "in-progress"
	DONE_STATUS       = "done"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewTask(description string) *Task {
	var id int

	tasks := LoadTasks()

	if len(tasks) == 0 {
		id = 1
	} else {
		id = len(tasks) + 1
	}

	task := Task{
		Id:          id,
		Description: description,
		Status:      TODO_STATUS,
		CreatedAt:   time.Now().Format(time.RFC822),
	}

	tasks = append(tasks, task)

	SaveTasks(tasks)

	return &task
}

func DeleteTask(id int) {

	tasks := LoadTasks()

	for index, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
		}
	}

	SaveTasks(tasks)
}

func ListTasks(args []string) {
	tasks := LoadTasks()

	if len(args) == 0 {
		fmt.Printf("%-3s | %-30s | %-10s\n", "ID", "Description", "Status")
		fmt.Println("---------------------------------------------")
		for _, task := range tasks {
			fmt.Printf("%-3d | %-30s | %-10s\n", task.Id, task.Description, task.Status)
		}
	} else if strings.ToLower(args[0]) == DONE_STATUS {
		fmt.Printf("%-3s | %-30s\n", "ID", "Description")
		fmt.Println("---------------------------------------------")
		for _, task := range tasks {
			if task.Status == DONE_STATUS {
				fmt.Printf("%-3d | %-30s\n", task.Id, task.Description)
			}
		}
	} else if strings.ToLower(args[0]) == TODO_STATUS {
		fmt.Printf("%-3s | %-30s\n", "ID", "Description")
		fmt.Println("---------------------------------------------")
		for _, task := range tasks {
			if task.Status == TODO_STATUS {
				fmt.Printf("%-3d | %-30s\n", task.Id, task.Description)
			}
		}
	} else if strings.ToLower(args[0]) == INPROGRESS_STATUS {
		fmt.Printf("%-3s | %-30s\n", "ID", "Description")
		fmt.Println("---------------------------------------------")
		for _, task := range tasks {
			if task.Status == INPROGRESS_STATUS {
				fmt.Printf("%-3d | %-30s\n", task.Id, task.Description)
			}
		}
	}
}

func Done(id int) {
	tasks := LoadTasks()

	for ind, task := range tasks {
		if task.Id == id {
			if task.Status == DONE_STATUS {
				fmt.Println("This issue is already done.")
				continue
			}
			tasks[ind].Status = DONE_STATUS
			tasks[ind].UpdatedAt = time.Now().Format(time.RFC822)
			fmt.Println(tasks[ind])
		} else {
			log.Println("error: task doesn't exist")
		}
	}

	SaveTasks(tasks)
}

func InProgress(id int) {

	tasks := LoadTasks()

	for ind, task := range tasks {
		if task.Id == id {
			if task.Status == INPROGRESS_STATUS {
				fmt.Println("This issue is already InProgress.")
				fmt.Println(tasks[ind])
				break
			}
			tasks[ind].Status = INPROGRESS_STATUS
			tasks[ind].UpdatedAt = time.Now().Format(time.RFC822)
			fmt.Println(tasks[ind])
		}
	}

	SaveTasks(tasks)
}

func Update(args []string) {
	tasks := LoadTasks()

	if len(args) < 2 {
		log.Println("error: not enough args to complete the command.")
		return
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println("error: first argument must be an task ID.")
		return
	}

	description := args[1]
	if description == "" {
		log.Println("error: description cannot be empty or must be int.")
		return
	}

	for ind, task := range tasks {
		if task.Id == taskID {
			tasks[ind].Description = args[1]
			tasks[ind].UpdatedAt = time.Now().Format(time.RFC822)
			fmt.Println(tasks[ind])
		}
	}

	SaveTasks(tasks)

}
