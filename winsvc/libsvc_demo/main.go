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
	log.Println("Loop logger boot, running...")

	tick := time.Tick(3 * time.Second)
	for {
		select {
		case <-p.sig:
			break
		case <-tick:
			cnt++
			log.Println("Loop logger count:", cnt)
		}
	}
}

func main() {

	logf, err := os.OpenFile("D:\\file.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logf)

	log.Println("Boot Args:", os.Args)

	name := "TestLoopLogger"
	config := &service.Config{
		Name:        name,
		DisplayName: "Go loop logger",
		Description: "test windows service， log file path: d:\\file.log",
	}
	pgm := &program{sig: make(chan bool)}
	s, err := service.New(pgm, config)
	if err != nil {
		log.Fatal(err)
	}
	//logger, err = s.Logger(nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	if len(os.Args) > 1 {
		// 接受控制命令 ["start", "stop", "restart", "install", "uninstall"]
		err := service.Control(s, os.Args[1])
		if err != nil {
			fmt.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}

		//switch os.Args[1] {
		//case "install":
		//	err := s.Install()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	log.Printf("服务 [%s] 安装成功", name)
		//	return
		//case "remove":
		//	err := s.Uninstall()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	log.Printf("服务 [%s] 卸载成功", name)
		//	return
		//case "start":
		//	err := s.Start()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	log.Printf("服务 [%s] 启动成功", name)
		//}
		return
	}

	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("服务停止")
}
