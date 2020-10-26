package main

import (
	"log"
	"os"
	"time"
)

func permission() {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		log.Println("Running with User permission.")
	} else {
		log.Println("Running with Admin permission.")

	}
}
func main() {
	file, err := os.OpenFile("D:/loop.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Boot Args:", os.Args)
	permission()
	cnt := 0
	for {
		cnt++
		log.Println("Log Cnt:", cnt)
		time.Sleep(time.Second * 3)
	}
}
