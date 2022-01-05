package db

import (
	"encoding/binary"
	"time"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB
var taskBucket = []byte("tasks")

type Task struct {
	ID   int
	Task string
}

func Init(dbPath string) error {
	var err error

	db, err = bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			return err
		}
		return nil
	})

}

// CreateTask creates a task and stores it into the database
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id := int(id64)
		return b.Put(itob(id), []byte(task))
	})
	if err != nil {
		return -1, err
	}

	return id, nil
}

// AllTasks returns the list of all tasks currently in the database
func AllTasks() ([]Task, error) {

	var tasks []Task
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				ID:   btoi(k),
				Task: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DelTask(key int) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})

}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64((b)))
}
