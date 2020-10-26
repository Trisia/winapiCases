package main

import (
	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/windows/elevate"
)

func main() {

	//path := "D:\\Project\\demo\\winapiCases\\winsvc\\install\\install.exe"
	path := "D:\\Project\\demo\\winapiCases\\winsvc\\uninstall\\uninstall.exe"
	// 管理员模式运行
	err := elevate.ShellExecute(path, "", "", windows.SW_SHOW)
	if err != nil {
		panic(err)
	}

}
