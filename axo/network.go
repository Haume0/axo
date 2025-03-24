package axo

import "net"

func HostIPs() []string {
	var links []string

	// Add localhost
	links = append(links, "localhost")

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return links
	}

	for _, iface := range interfaces {
		// Skip loopback and interfaces that are down
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			// Check if it's an IP network address
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				links = append(links, ipnet.IP.String())
			}
		}
	}

	return links
}
