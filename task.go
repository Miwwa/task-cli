package main

import "time"

type TaskId uint
type TaskStatus string

const (
	Todo       = "todo"
	InProgress = "in-progress"
	Done       = "done"
)

type Task struct {
	Id          TaskId     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
