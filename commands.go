package main

import "fmt"

const helpText = `Usage: task-cli <command> [arguments]

task-cli add <description>			add new task with specified <description>
task-cli update <id> <description>  update description of task with specified <id>
task-cli delete <id>				delete task by specified <id>
task-cli mark-in-progress <id>		mark in-progress task by specified <id>
task-cli mark-done <id>				mark done task specified <id>
task-cli list						list all tasks
task-cli list <status>				list tasks by status (done, todo, in-progress)
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

	switch args[0] {
	case "add":
		return AddCmd(args[1:])
	case "update":
		return UpdateCmd(args[1:])
	case "delete":
	case "mark-in-progress":
	case "mark-done":
	case "list":
	}
	return HelpCmd()
}

func HelpCmd() (string, error) {
	return helpText, nil
}

func AddCmd(args []string) (string, error) {
	if len(args) == 0 || len(args) > 1 {
		return "", InvalidUsageError("add <description>")
	}
	return fmt.Sprintf("Task added successfully (ID: %s)", args[0]), nil
}

func UpdateCmd(args []string) (string, error) {
	if len(args) == 0 || len(args) > 2 {
		return "", InvalidUsageError("update <id> <description>")
	}
	return fmt.Sprintf("Task updated successfully (ID: %s)", args[0]), nil
}
