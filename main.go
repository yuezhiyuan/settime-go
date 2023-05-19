package main

import (
	"log"
	"time"
)

func main() {
	t, _ := time.Parse(time.DateTime, time.DateTime)
	log.Println("time:", time.DateTime)
	err := SetTime(t)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("time:", time.Now().Format(time.DateTime))
	time.Sleep(15 * time.Second)
}
