package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const helpText = `Usage: task-cli <command> [arguments]

task-cli add <description>              add new task with specified <description>
task-cli update <id> <description>      update description of task with specified <id>
task-cli delete <id>                    delete task by specified <id>
task-cli mark-in-progress <id>          mark in-progress task by specified <id>
task-cli mark-done <id>                 mark done task specified <id>
task-cli list                           list all tasks
task-cli list <status>                  list tasks by status (done, todo, in-progress)
`

type CmdError struct {
	message string
	usage   string
}

func (e *CmdError) Error() string {
	return fmt.Sprintf("%s\nUsage: task-cli %s", e.message, e.usage)
}

func InvalidUsageError(usage string) error {
	return &CmdError{message: "invalid arguments", usage: usage}
}

func Run(args []string) (string, error) {
	if len(args) == 0 {
		return HelpCmd()
	}

	tasks, err := FromFile("./tasks.json")
	if err != nil {
		return "", err
	}

	switch args[0] {
	case "add":
		return AddCmd(args[1:], tasks)
	case "update":
		return UpdateDescriptionCmd(args[1:], tasks)
	case "delete":
		return DeleteCmd(args[1:], tasks)
	case "mark-in-progress":
		return UpdateStatusCmd(args[1:], tasks, InProgress)
	case "mark-done":
		return UpdateStatusCmd(args[1:], tasks, Done)
	case "list":
		return ListCmd(args[1:], tasks)
	}
	return HelpCmd()
}

func HelpCmd() (string, error) {
	return helpText, nil
}

func AddCmd(args []string, storage *TasksStorage) (string, error) {
	if len(args) == 0 || len(args) > 1 {
		return "", InvalidUsageError("add <description>")
	}
	task, err := storage.Add(args[0])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task added successfully (ID: %d)", task.Id), nil
}

func UpdateDescriptionCmd(args []string, storage *TasksStorage) (string, error) {
	if len(args) == 0 || len(args) > 2 {
		return "", InvalidUsageError("update <id> <description>")
	}
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return "", err
	}
	task, err := storage.UpdateDescription(TaskId(id), args[1])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task updated successfully (ID: %d)", task.Id), nil
}

func UpdateStatusCmd(args []string, storage *TasksStorage, status TaskStatus) (string, error) {
	if len(args) == 0 || len(args) > 1 {
		return "", InvalidUsageError(fmt.Sprintf("%s <id>", args[0]))
	}
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return "", err
	}
	task, err := storage.UpdateStatus(TaskId(id), status)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task updated successfully (ID: %d)", task.Id), nil
}

func DeleteCmd(args []string, storage *TasksStorage) (string, error) {
	if len(args) == 0 || len(args) > 1 {
		return "", InvalidUsageError("delete <id>")
	}
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return "", err
	}
	err = storage.Delete(TaskId(id))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task deleted successfully (ID: %s)", args[0]), nil
}

func ListCmd(args []string, storage *TasksStorage) (string, error) {
	var listUsage = fmt.Sprintf("list [%s|%s|%s]", Todo, InProgress, Done)
	if len(args) > 1 {
		return "", InvalidUsageError(listUsage)
	}
	if len(args) == 0 {
		return formatTasks(storage.GetAll()), nil
	}
	switch args[0] {
	case Todo, InProgress, Done:
		return formatTasks(storage.GetByStatus(TaskStatus(args[0]))), nil
	}
	return "", InvalidUsageError(listUsage)
}

func formatTasks(tasks []Task) string {
	var sb strings.Builder
	for _, task := range tasks {
		sb.WriteString(fmt.Sprintf(
			"ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			task.Id, task.Description, task.Status, task.CreatedAt.Format(time.RFC822), task.UpdatedAt.Format(time.RFC822),
		))
	}
	return sb.String()
}
