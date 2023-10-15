package manager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Manager struct {
	id     uint8
	tasks  map[uint8]string
	reader *bufio.Reader
}

func NewManager() Manager {
	return Manager{
		1,
		make(map[uint8]string),
		bufio.NewReader(os.Stdin),
	}
}

// TODO: Add an option to delete a task from the list
// TODO: Add an option to save the tasks for a file
// TODO: After adding an option to save to file, consider adding and removing from a file directly.

func (m *Manager) CreateNewTask() {
	fmt.Print("Enter task description: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	m.tasks[m.id] = input
	m.id++
}

func (m *Manager) EditTask() error {
	m.ViewTasks()
  taskid, err := m.userInputToTaskID("Choose task by id: ")
  if err != nil {
    return err
  }
	for {
    _, exists := m.tasks[taskid]
    if exists {
      break
    } 
    taskid, err = m.userInputToTaskID("Please enter a valid task id: ")
    if err != nil {
      return err
    }
	}
  fmt.Print("Enter new description: ")
  newDescription, _ := m.reader.ReadString('\n')
  m.tasks[taskid] = newDescription
	return nil
}

func (m *Manager) DeleteTask() error {
	m.ViewTasks()
	taskid, err := m.userInputToTaskID("Task id: ")
	if err != nil {
		return err
	}
	delete(m.tasks, taskid)
	return nil
}

func (m *Manager) ViewTasks() {
	for index := uint8(1); index < m.id ; index++ {
    _, exists := m.tasks[index]
    if exists {
		  fmt.Printf("%v. %s", index, m.tasks[index])
    }
	}
}

// Convert the user input to uint8.
func (m *Manager) userInputToTaskID(msg string) (uint8, error) {
	fmt.Print(msg)
	input, _ := m.reader.ReadString('\n')
	input = strings.TrimSpace(input)
	conv, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return 0, errors.New("a problem occourd")
	}
	if conv > 255 {
		return 0, errors.New("encouterd a problem with task id.")
	}
	return uint8(conv), nil
}
