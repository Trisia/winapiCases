package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
	"time"
)

type program struct {
}

func (p *program) Start(s service.Service) error {
	return nil
}

func (p *program) Stop(s service.Service) error {
	// 退出时删除服务注册
	err := s.Uninstall()
	return err
}

func main() {

	serviceName := "ForceInstallSvc"
	cfg := &service.Config{
		Name:        serviceName,
		DisplayName: "ForceInstallSvc",
		Description: "强制安装测试，如果已经存在那么会删除过去的服务",
		Arguments:   []string{"Service"},
	}
	pgm := &program{}
	s, err := service.New(pgm, cfg)
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 && os.Args[1] == "Service" {
		_ = s.Run()
		return
	}

	// 检查是否已经安装
	status, err := s.Status()
	if err == nil {
		if status == service.StatusRunning {
			log.Println("服务正在运行，停止服务")
			// 如果服务正在运行那么首先停止服务
			_ = s.Stop()
		}
		// 卸载服务
		log.Println("卸载服务")
		err = s.Uninstall()
		if err != nil {
			panic(err)
		}
		// 等待服务卸载完成
		for {
			_, err = s.Status()
			if err != nil {
				break
			}
			time.Sleep(time.Second / 3)
		}
	} else {
		if err != service.ErrNotInstalled {
			panic(err)
		}
	}

	// 安装服务
	log.Println("安装服务")
	err = s.Install()
	if err != nil {
		panic(err)
	}
	log.Println("启动服务")
	err = s.Start()
	if err != nil {
		panic(err)
	}
}
