package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
)

func main() {
	m, err := mgr.Connect()
	if err != nil {
		panic(err)
	}
	defer m.Disconnect()
	svcName := "TestLoopLogService"

	svc, err := m.OpenService(svcName)
	if err != nil {
		log.Fatal(err)
	}
	defer svc.Close()

	err = svc.Delete()
	if err != nil {
		log.Fatalf("Service Delete failed: %v", err)
	}

	log.Printf("Service [%s] deleted.", svcName)
	log.Println("Press Enter to finish.")
	fmt.Scanln()
}
