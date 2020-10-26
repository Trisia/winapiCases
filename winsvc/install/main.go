package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"time"
)

func main() {
	svcName := "TestLoopLogService"

	m, err := mgr.Connect()
	if err != nil {
		panic(err)
	}
	defer m.Disconnect()

	c := mgr.Config{
		StartType:   mgr.StartManual,
		DisplayName: "TestLoopLogService",
		Description: "This Service is test service, func: loop run log to file D:\\loop.log",
	}
	exepath := "D:\\Project\\demo\\winapiCases\\winsvc\\demo_svc\\demo_svc.exe"

	install(m, svcName, exepath, c)
	log.Printf("Service [%s] installed", svcName)
	fmt.Println("Press Enter to finish.")
	fmt.Scanln()

}

func install(m *mgr.Mgr, name string, exepath string, c mgr.Config) {
	svc, err := m.OpenService(name)
	// 等待一会获取服务
	for i := 0; ; i++ {
		if err != nil {
			break
		}
		svc.Close()
		if i > 10 {
			log.Fatalf("Service %s already exists", name)
		}
		time.Sleep(300 * time.Millisecond)
	}

	svc, err = m.CreateService(name, exepath, c)
	if err != nil {
		log.Fatalf("Create Service(%s) failed: %v", name, err)
	}
	defer svc.Close()
}
