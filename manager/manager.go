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
	id   uint8
	tasks map[uint8]string
}

func NewManager() Manager {
  return Manager{
    1,
    make(map[uint8]string),
  }
}

func (m *Manager) CreateNewTask() {
  fmt.Print("Enter task description: ")
	reader := bufio.NewReader(os.Stdin)
  input,_ := reader.ReadString('\n')
  m.tasks[m.id] = input
  m.id++
}

func (m *Manager) EditTask() error {
  m.ViewTasks()
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Choose task by id: ")
  taskid,_ := reader.ReadString('\n')
  taskid = strings.TrimSpace(taskid)
  // converting the input to uint8
  conv,err := strconv.ParseUint(taskid, 10, 8)
  if err != nil {
    fmt.Println("Cannot get the input")
    return errors.New("a problem occourd")
  }
  if conv > 255 {
    return errors.New("encouterd a problem with task id.")
  }
  id := uint8(conv)
  // ------------------------------
  fmt.Print("Enter new description: ")
  newDescription,_ := reader.ReadString('\n')
  m.tasks[id] = newDescription
  return nil
}

func (m *Manager) ViewTasks() {
  for index, task := range m.tasks {
    fmt.Printf("%v. %s", index,task)
  }
}
