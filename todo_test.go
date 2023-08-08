package todo_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/fahadqazi/golang/todo-app"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, but got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New task"
	l.Add(taskName)

	if l[0].Done {
		t.Errorf("New task should not be complete")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}

	greaterThanLen := l.Complete(5)

	comparisonOne := greaterThanLen.Error() == fmt.Errorf("item 5 does not exist").Error()

	if !comparisonOne {
		t.Errorf("Did not return correct error message")
	}

	lessThanOne := l.Complete(-1)

	comparisonTwo := lessThanOne.Error() == fmt.Errorf("item -1 does not exist").Error()

	if !comparisonTwo {
		t.Errorf("Did not return correct error message")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	for _, t := range tasks {
		l.Add(t)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, but got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length to be %d, but got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected list length to be %q, but got %q", tasks[2], l[1].Task)
	}

	greaterThanLen := l.Delete(5)

	comparisonOne := greaterThanLen.Error() == fmt.Errorf("item 5 does not exist").Error()

	if !comparisonOne {
		t.Errorf("Did not return correct error message")
	}

	lessThanOne := l.Delete(-1)

	comparisonTwo := lessThanOne.Error() == fmt.Errorf("item -1 does not exist").Error()

	if !comparisonTwo {
		t.Errorf("Did not return correct error message")
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, but got %q instead", taskName, l1[0].Task)
	}

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Errorf("Error creating temp file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	if err := l1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tempFile.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}

func TestSave(t *testing.T) {
	t.Run("Successful save", func(t *testing.T) {
		list := &todo.List{}

		tempFile, err := ioutil.TempFile("", "testlist.json")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		err = list.Save(tempFile.Name())
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
