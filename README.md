# Task Tracker CLI

Made for https://roadmap.sh/projects/task-tracker

Task Tracker CLI is a simple command-line application to manage your tasks. You can add, update, delete,
and change the status of your tasks, as well as list all tasks or filter them based on their status.

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed on your system.
2. Clone this repository to your local machine.
3. Navigate to the project directory.
4. Run `go build` to compile the application.

```sh
git clone https://github.com/Miwwa/task-cli
cd task-tracker-cli
go build
```

## Usage

The CLI application accepts various commands with corresponding arguments.

```
Usage: task-cli <command> [arguments]

task-cli add <description>              add new task with specified <description>
task-cli update <id> <description>      update description of task with specified <id>
task-cli delete <id>                    delete task by specified <id>
task-cli mark-in-progress <id>          mark in-progress task by specified <id>
task-cli mark-done <id>                 mark done task specified <id>
task-cli list                           list all tasks
task-cli list <status>                  list tasks by status (done, todo, in-progress)
```
