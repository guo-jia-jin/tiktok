package initialize

import (
	"fmt"
	"net"
	"strings"

	"Tiktok/global"
)

func getOutBoundIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func GetStaticUrl() (urlhead string) {
	ip := getOutBoundIP()
	urlhead = "http://" + ip + global.GVA_CONFIG.RunPort.Content
	return
}
