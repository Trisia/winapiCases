package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/windows/tunnel/winipcfg"
	"log"
	"net"
)

// 创建端口
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

// 配置端口IP
func setInterfaceCfg(device tun.Device) error {
	nativeTunDevice := device.(*tun.NativeTun)
	luid := winipcfg.LUID(nativeTunDevice.LUID())
	ip, ipNet, _ := net.ParseCIDR("10.0.0.21/24")
	err := luid.SetIPAddresses([]net.IPNet{{ip, ipNet.Mask}})
	if err != nil {
		return err
	}
	targetHostIp, targetIpNet, _ := net.ParseCIDR("172.31.214.12/32")
	targetIpNet.IP = targetHostIp

	err = luid.SetRoutes([]*winipcfg.RouteData{
		{*ipNet, ipNet.IP, 0},
		{*targetIpNet, ip, 0},
	})
	if err != nil {
		return err
	}
	return nil
}
