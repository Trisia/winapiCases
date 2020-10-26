package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/tun"
	"log"
)

func tunInterfaceCreate() (tun.Device, error) {
	infName := "GoWinTun"
	deivce, err := tun.CreateTUN(infName, 0)
	//tun, err := tun.CreateTUNWithRequestedGUID(infName, "",0)
	if err == nil {
		realName, err := deivce.Name()
		if err == nil {
			infName = realName
			log.Printf("Wintun设备 [%s] 创建成功", realName)
		}
	} else {
		return nil, fmt.Errorf("failed to tunInterfaceCreate TUN device: %v", err)
	}
	return deivce, nil
}
