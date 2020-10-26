package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/tun"
)

func tunInterfaceCreate() (tun.Device, error) {
	infName := "Hell"
	tun, err := tun.CreateTUN(infName, 0)
	if err == nil {
		realName, err := tun.Name()
		if err == nil {
			infName = realName
		}
	} else {
		return nil, fmt.Errorf("failed to tunInterfaceCreate TUN device: %v", err)
	}
	return tun, nil
}
