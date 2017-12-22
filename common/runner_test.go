package common

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {

	log.Println("...开始执行任务...")
	timeout := 3 * time.Second
	r := NewRunner(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}
func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
