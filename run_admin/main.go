package main

import (
	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/windows/elevate"
	"os/exec"
)

func main() {

	path := "D:\\Project\\demo\\winapiCases\\run_admin\\PermissionCase.exe"
	// 管理员模式运行
	err := elevate.ShellExecute(path, "", "", windows.SW_SHOW)
	if err != nil {
		panic(err)
	}

	// 用户模式运行
	c := exec.Command("explorer", path)
	if err := c.Run(); err != nil {
		panic(err)
	}
	return
}
