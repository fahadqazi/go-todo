package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) String() string {
	headerColor := color.New(color.FgCyan).Add(color.Underline)
	header := headerColor.Sprintf("%s", "TODO APP")
	formatted := fmt.Sprintf("----- %s -----\n", header)
	doneColor := color.New(color.FgRed).Add(color.CrossedOut).Add(color.Faint)
	todoColor := color.New(color.FgGreen).Add(color.Bold)

	for k, t := range *l {
		formattedTime := fmt.Sprintf(t.CreatedAt.Format("Mon Jan 2 15:04:05"))
		postFix := ""
		if t.Done {
			postFix = " ✔︎"
			t.Task = doneColor.Sprintf("%s%s", t.Task, postFix)
		} else {
			t.Task = todoColor.Sprintf("%s", t.Task)
		}

		formatted += fmt.Sprintf("%d: [%s] %s\n", k+1, formattedTime, t.Task)
	}

	return formatted
}

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *List) Clear() {
	ls := *l
	ls = nil
	*l = ls
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

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
