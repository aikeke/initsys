package main

import (
	"fmt"
	"initsys"
	//"os/exec"
	//"os"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		eth = kingpin.Flag(
			"nic",
			"需要修改的网卡名称",
		).Default("eno1").String()
		ip = kingpin.Flag(
			"ip",
			"主机ip地址",
		).String()
		ssh = kingpin.Flag(
			"ssh.enable",
			"是否优化ssh配置",
		).Default("true").Bool()
		hostname = kingpin.Flag(
			"hostname",
			"需要修改的主机名",
		).String()
	)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	fmt.Println(*eth,*ip,*ssh,*hostname)
	s:=initsys.NetInit(*eth,*ip)
	fmt.Println(s)

}
