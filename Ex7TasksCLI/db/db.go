package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB
var bucketName = "TasksBucker"

type Task struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Done  bool   `json:"done"`
}

func Init(dbpath string) {
	var err error
	db, err = bolt.Open(dbpath, 0666, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	db.Update(func(tx *bolt.Tx) error {
		// Use the transaction...
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

func CreateTask(taskString string) (Task, error) {
	var newTask Task
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		curTime := time.Now().Format(time.RFC3339)
		newTask = Task{
			Key:   curTime,
			Value: taskString,
			Done:  false,
		}
		bytes, _ := json.Marshal(newTask)
		err := b.Put([]byte(curTime), bytes)
		return err
	})
	return newTask, err
}

func AllTasks() ([]Task, error) {
	var list []Task
	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t Task
			_ = json.Unmarshal(v, &t)
			list = append(list, t)
		}

		return nil
	})
	return list, err
}

func DoTask(key string) (Task, error) {
	var t Task
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		taskBytes := b.Get([]byte(key))
		_ = json.Unmarshal(taskBytes, &t)
		t.Done = true
		taskBytes, _ = json.Marshal(t)
		return b.Put([]byte(key), taskBytes)
	})
	return t, err
}

func IncompleteTask() ([]Task, error) {
	all, err := AllTasks()
	if err != nil {
		return nil, err
	}
	var done []Task
	for _, t := range all {
		if !t.Done {
			done = append(done, t)
		}
	}
	return done, nil
}

func DeleteTask(key string) (Task, error) {
	var t Task
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		taskBytes := b.Get([]byte(key))
		_ = json.Unmarshal(taskBytes, &t)
		return b.Delete([]byte(key))
	})
	return t, err
}

func TasksInTimeRange(minTime, maxTime string) ([]Task, error) {
	var list []Task
	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		min := []byte(minTime)
		max := []byte(maxTime)

		// Iterate over the 90's.
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			var t Task
			_ = json.Unmarshal(v, &t)
			list = append(list, t)
		}
		return nil
	})
	return list, err
}

func CompletedTasks(minTime, maxTime string) ([]Task, error) {
	list, err := TasksInTimeRange(minTime, maxTime)
	if err != nil {
		return nil, err
	}
	var ret []Task
	for _, task := range list {
		if task.Done {
			ret = append(ret, task)
		}
	}
	return ret, nil
}
