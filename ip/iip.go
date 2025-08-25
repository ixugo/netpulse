package ip

import (
	"net"
	"time"
)

func InternalIP() string {
	conn, err := net.DialTimeout("udp", "8.8.8.8:53", 3*time.Second)
	if err != nil {
		return ""
	}
	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)
	if addr.IP.To4() != nil {
		return addr.IP.String()
	}

	return localIP()
}

// localIP 获取本地 IP，遇到虚拟 IP 有概率不准确
func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, v := range addrs {
		net, ok := v.(*net.IPNet)
		if !ok {
			continue
		}
		if net.IP.IsMulticast() || net.IP.IsLoopback() || net.IP.IsLinkLocalMulticast() || net.IP.IsLinkLocalUnicast() {
			continue
		}
		if ip := net.IP.To4(); ip != nil {
			return ip.String()
		}
	}
	return ""
}
