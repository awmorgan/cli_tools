package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item represents a Todo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of Todo items
type List []item

// Add creates a new todo item and adds it to the list
func (l *List) Add(task string) {
	*l = append(*l, item{Task: task, CreatedAt: time.Now()})
}

// Complete marks the task as done
func (l *List) Complete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("item %d does not exist", i)
	}
	(*l)[i-1].Done = true
	(*l)[i-1].CompletedAt = time.Now()
	return nil
}

// Delete removes the item from the list
func (l *List) Delete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = append((*l)[:i-1], (*l)[i:]...)
	return nil
}

// Save encode the list to JSON and writes it to a file
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get retrives a list from a JSON file
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
