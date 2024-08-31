package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type TasksStorage struct {
	filePath string
	NextId   TaskId          `json:"nextId"`
	Tasks    map[TaskId]Task `json:"tasks"`
}

func FromFile(path string) (*TasksStorage, error) {
	storage := &TasksStorage{filePath: path, NextId: 1, Tasks: map[TaskId]Task{}}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return storage, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return &TasksStorage{}, err
	}
	err = json.Unmarshal(data, &storage)
	if err != nil {
		return &TasksStorage{}, err
	}
	return storage, nil
}

func (storage *TasksStorage) Save() error {
	data, err := json.MarshalIndent(storage, "", "   ")
	if err != nil {
		return err
	}
	err = os.WriteFile(storage.filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (storage *TasksStorage) Add(description string) (Task, error) {
	task := Task{
		Id:          storage.NextId,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	storage.Tasks[task.Id] = task
	storage.NextId++
	err := storage.Save()
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (storage *TasksStorage) UpdateDescription(id TaskId, newDescription string) (Task, error) {
	task, ok := storage.Tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("task with id %d not found", id)
	}
	task.Description = newDescription
	task.UpdatedAt = time.Now()
	storage.Tasks[id] = task
	err := storage.Save()
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (storage *TasksStorage) UpdateStatus(id TaskId, newStatus TaskStatus) (Task, error) {
	task, ok := storage.Tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("task with id %d not found", id)
	}
	task.Status = newStatus
	task.UpdatedAt = time.Now()
	storage.Tasks[id] = task
	err := storage.Save()
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (storage *TasksStorage) Delete(id TaskId) error {
	delete(storage.Tasks, id)
	err := storage.Save()
	if err != nil {
		return err
	}
	return nil
}

func (storage *TasksStorage) GetAll() []Task {
	tasks := make([]Task, 0, len(storage.Tasks))
	for _, task := range storage.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (storage *TasksStorage) GetByStatus(status TaskStatus) []Task {
	tasks := make([]Task, 0, len(storage.Tasks))
	for _, task := range storage.Tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
