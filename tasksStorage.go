package main

import (
	"encoding/json"
	"errors"
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
