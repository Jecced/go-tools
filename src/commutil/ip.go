package commutil

import (
	"errors"
	"github.com/Jecced/go-tools/src/https"
	"net"
)

// 获取本机内网ip
func GetInternal() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String(), nil
					}
				}
			}
		}
	}

	return "", errors.New("net.Interfaces failed")
}

// 获取本机公网ip
func GetExternal() (string, error) {
	return https.Get("https://myexternalip.com/raw").Send().ReadText()
}
