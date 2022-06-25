package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []item

func (t *TodoList) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *TodoList) Complete(index int) error {
	ls := *t
	if index < 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *TodoList) Delete(index int) error {
	ls := *t
	if index < 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *TodoList) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoList) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func (t *TodoList) List() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell
	for i, v := range *t {
		i += 1
		task := blue(v.Task)
		done := blue("no")
		createAt := v.CreatedAt.Format(time.RFC822)
		completedAt := v.CompletedAt.Format(time.RFC822)
		if v.Done {
			task = green(fmt.Sprintf("\u2705 %s", v.Task))
			done = green("yes")
			createAt = green(createAt)
			completedAt = green(completedAt)
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: task},
			{Text: fmt.Sprintf("%s", done)},
			{Text: createAt},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{
		Cells: cells,
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("you have [%d] pending todos!", t.countPendingTodos()))},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (t *TodoList) countPendingTodos() int {
	total := 0
	for _, v := range *t {
		if !v.Done {
			total++
		}
	}
	return total
}
