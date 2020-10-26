package main

import (
	"golang.zx2c4.com/wireguard/tun"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func tunInterfaceCreate() {
	infName := "Hell"
	tun, err := tun.CreateTUN(infName, 0)
	if err == nil {
		realName, err := tun.Name()
		if err == nil {
			infName = realName
		}
	} else {
		log.Fatalln("Failed to tunInterfaceCreate TUN device:", err)
	}

	log.Printf("Tun Device [%s] tunInterfaceCreate success.", infName)

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt)
	signal.Notify(term, os.Kill)
	signal.Notify(term, syscall.SIGTERM)

	select {
	case <-term:
	}
	log.Println("Service down.")
}
