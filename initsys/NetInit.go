package initsys

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func NetInit(eth string, ip string) string {
	filepath := "/etc/sysconfig/network-scripts/ifcfg-" + eth
	_, err := PathExists(filepath)
	if err != nil {
		fmt.Println(err)
		return "failed"
	}
	addr := net.ParseIP(ip)
	if addr == nil {
		fmt.Println("ip地址错误")
		return "failed"
	}
	err = os.Rename(filepath, filepath+"_bak")
	if err != nil {
		fmt.Println(err)
		return "failed"
	}
	srcfile, srcerr := os.Open(filepath + "_bak")
	if srcerr != nil {
		fmt.Println(srcerr)
		return "failed"
	}
	defer srcfile.Close()
	srcreader := bufio.NewReader(srcfile)
	destfile, desterr := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if desterr != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return "failed"
	}
	defer destfile.Close()
	destwriter := bufio.NewWriter(destfile)
	flag := 0
	area := strings.Split(ip, ".")[1]
	for {
		srcstring, readerr := srcreader.ReadString('\n')
		if readerr == io.EOF {
			break
		}
		if s := strings.Contains(srcstring, "ONBOOT"); s {

			srcstring = "ONBOOT=yes\n"
		}
		if s := strings.Contains(srcstring, "BOOTPROTO"); s {

			srcstring = "BOOTPROTO=static\n"
		}
		if s := strings.Contains(srcstring, "IPADDR"); s {

			srcstring = "IPADDR=" + ip + "\n"
			flag = 1
		}
		if s := strings.Contains(srcstring, "NETMASK"); s {

			srcstring = "NETMASK=255.255.0.0\n"
		}
		if s := strings.Contains(srcstring, "GATEWAY"); s {

			srcstring = "GATEWAY=172." + area + ".255.254\n"
		}
		destwriter.WriteString(srcstring)

	}
	if flag == 0 {
		destwriter.WriteString("NETMASK=255.255.0.0\n" + "IPADDR=" + ip + "\n" + "GATEWAY=172." + area + ".255.254\n")
	}
	destwriter.Flush()
	err = os.Remove(filepath + "_bak")
	if err != nil {
		fmt.Println(err)
		return "failed"
	}
	return "success"

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}
