package main

import (
	// "fmt"
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Start() {
	job.Println("Starting now....")
	job.Println("end")
}

func main() {
	mylog := log.New(os.Stdout, "", log.LstdFlags)
	j := NewJob("test", mylog)
	j.Start()
}