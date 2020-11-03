package main

import (
	"fmt"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	device2 "golang.zx2c4.com/wireguard/device"
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
	//targetHostIp, targetIpNet, _ := net.ParseCIDR("172.31.214.12/32")
	//targetIpNet.IP = targetHostIp

	err = luid.SetRoutes([]*winipcfg.RouteData{
		{*ipNet, net.ParseIP("10.0.0.1"), 0},
		//{*targetIpNet, ip, 0},
	})
	if err != nil {
		return err
	}
	return nil
}

// 从TUN设备中读取数据并打印日志
func readFromTUNLog(dev tun.Device) {
	for {
		buffer := make([]byte, 65535)
		n, err := dev.Read(buffer[:], 0)
		if err != nil {
			log.Println("Failed to read packet from TUN device:", err)
			return
		}
		pkg := buffer[0:n]
		if n == 0 || n > device2.MaxContentSize {
			return
		}
		switch pkg[0] >> 4 {
		case ipv4.Version:
			if len(pkg) < ipv4.HeaderLen {
				return
			}
			h, err := ipv4.ParseHeader(pkg)
			if err != nil {
				log.Println("Error parse ipv4 header:", err)
				return
			}
			log.Printf(">> Protocol [0x%02X] src: %s, dst: %s", h.Protocol, h.Src.String(), h.Dst.String())
		case ipv6.Version:
			if len(pkg) < ipv6.HeaderLen {
				return
			}
			h, err := ipv4.ParseHeader(pkg)
			if err != nil {
				log.Println("Error parse ipv4 header:", err)
				return
			}
			log.Printf(">> Protocol [0x%02X] src: %s, dst: %s", h.Protocol, h.Src.String(), h.Dst.String())
		}
	}
}
