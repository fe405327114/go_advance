package main

import (
	"time"
	"log"
	"os"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work")
	r :=runner.New(timeout)
	r.Add(createTask(),createTask(),createTask())
	if err:=r.Start();err!=nil{
		switch err {
		case runner.ErrTimeout:
			log.Println("timeout")
		os.Exit(1)
		case runner.Errinterrupt:
			log.Println("interrupt")
		os.Exit(2)
		}
	}
	log.Println("Process End")
}
func createTask()func(int){
	return func(id int) {
		log.Printf("Processor-Task #%d",id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}