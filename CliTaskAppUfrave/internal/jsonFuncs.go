package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (t Task) String() string {
	return fmt.Sprintf("Id: %v\nDescription: %v\nStatus: %v\nCreatedAt: %v\nUpdatedAt: %v", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
}

func LoadTasks() (tasks []Task) {
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatalf("Error openning file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	return
}

func SaveTasks(tasks []Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}
