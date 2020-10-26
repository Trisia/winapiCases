package main

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
	"time"
)

var logger service.Logger

type program struct {
	sig chan bool
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	p.sig <- true
	return nil
}
func (p *program) run() {
	cnt := 0
	log.Println("测试程序启动")

	tunInterfaceCreate()

	tick := time.Tick(3 * time.Second)
	for {
		select {
		case <-p.sig:
			break
		case <-tick:
			cnt++
			log.Println("status: Alive")
		}
	}
}

func main() {

	logf, err := os.OpenFile("D:\\wintun.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logf)

	log.Println("Boot Args:", os.Args)

	name := "TestTunCases"
	config := &service.Config{
		Name:        name,
		DisplayName: "WinTunTestCases",
		Description: "test windows service， log file path: D:\\wintun.log",
	}
	pgm := &program{sig: make(chan bool)}
	s, err := service.New(pgm, config)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		// 接受控制命令 ["start", "stop", "restart", "install", "uninstall"]
		err := service.Control(s, os.Args[1])
		if err != nil {
			fmt.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}

	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("服务停止")
}
