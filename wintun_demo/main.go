package main

import (
	"fmt"
	"github.com/kardianos/service"
	"golang.zx2c4.com/wireguard/tun"
	"log"
	"os"
)

var logger service.Logger
var device tun.Device

type program struct {
	sig chan bool
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	p.sig <- true
	if device != nil {
		name, _ := device.Name()
		log.Printf("删除 [%s] wintun 设备", name)
		device.Close()
	}
	return nil
}
func (p *program) run() {
	log.Println("测试程序启动")
	var err error
	log.Printf("创建 网络适配器")
	device, err = tunInterfaceCreate()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("配置IP 路由")
	err = setInterfaceCfg(device)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("启动 TUN读取goroutine")
	//cnt := 0
	//tick := time.Tick(3 * time.Second)
	// 读取
	go readFromTUNLog(device)
	for {
		select {
		case <-p.sig:
			break
			//case <-tick:
			//	cnt++
			//	log.Println("status: Alive")
		}
	}
}

func main() {

	//logf, err := os.OpenFile("C:\\wintun.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 755)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.SetOutput(logf)

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
