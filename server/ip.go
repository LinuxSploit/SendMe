package server

import "net"

// GetLocalIP retrieves the local IPv4 address of the machine.
func GetServerLocalIP() string {
	// Get a list of network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return "0.0.0.0"
	}

	// Iterate through the network interfaces
	for _, iface := range interfaces {
		// Check if the interface is not a loopback and is up
		if isPhysicalNetworkInterface(iface) {
			// Get the addresses for this interface
			addrs, err := iface.Addrs()
			if err != nil {
				return "0.0.0.0"
			}

			// Iterate through the addresses and look for an IPv4 address
			for _, addr := range addrs {
				if ip, isIPv4 := getIPv4Address(addr); isIPv4 {
					return ip
				}
			}
		}
	}

	return "0.0.0.0"
}

// isPhysicalNetworkInterface checks if the network interface is not a loopback and is up.
func isPhysicalNetworkInterface(iface net.Interface) bool {
	return iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0
}

// getIPv4Address extracts the IPv4 address from the network address.
func getIPv4Address(addr net.Addr) (string, bool) {
	if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		if ip4 := ipnet.IP.To4(); ip4 != nil {
			return ip4.String(), true
		}
	}
	return "", false
}
