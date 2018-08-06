package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	// db is same as the one on the global scope
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		return bucket.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	/*
		Fetch all tasks
	*/
	var tasks []Task
	// use view function for READing data
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		return bucket.Delete(itob(key))
	})
}

func itob(v int) []byte {
	/*
		Integer to binary
	*/
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
