package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func main() {
	serviceName := "TestLoopLogger"

	m, err := mgr.Connect()
	if err != nil {
		panic(err)
	}

	service, err := m.OpenService(serviceName)
	if err != nil {
		// panic: The specified service does not exist as an installed service.
		panic(err)
	}
	defer service.Close()

	status, err := service.Query()
	if err != nil {
		panic(err)
	}
	switch status.State {
	case svc.Stopped:
		fmt.Println("Stopped")
	case svc.StartPending:
		fmt.Println("StartPending")
	case svc.StopPending:
		fmt.Println("StopPending")
	case svc.Running:
		fmt.Println("Running")
	case svc.ContinuePending:
		fmt.Println("ContinuePending")
	case svc.PausePending:
		fmt.Println("PausePending")
	case svc.Paused:
		fmt.Println("Pause")
	}

}
