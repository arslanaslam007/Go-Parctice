package main

import (
	"fmt"
	"test-go/kafka"
	"time"
)

func main() {

	fmt.Println("Starting...")
	go kafka.StartKafka()
	fmt.Println("kafka has been started...")
	time.Sleep(10 * time.Minute)
}
