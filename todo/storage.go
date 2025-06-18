package todo

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	taskFilePath string
	once sync.Once
)

func TaskFilePath() string {
	once.Do(func() {
		home, err := os.UserHomeDir()
		if err != nil {
			panic("couldn't get user home directory: " + err.Error())
		}

		taskFilePath = filepath.Join(home, ".tasks.json")
	})

	return taskFilePath
}