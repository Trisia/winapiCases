package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"time"
)

func main() {
	for {
		//process, err := os.FindProcess(5548)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println("Running ",process.Pid)
		//os.ProcessState{}

		process, err := ps.FindProcess(5548)
		if err != nil {
			panic(err)
		}
		if process == nil {
			fmt.Println("Not Running")
			break
		}
		fmt.Println("Running ", process.Pid())
		time.Sleep(time.Second)
	}
}
