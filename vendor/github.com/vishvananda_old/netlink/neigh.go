package netlink

import (
	"fmt"
	"net"
)

// Neigh represents a link layer neighbor from netlink.
type Neigh struct {
	LinkIndex    int
	Family       int
	State        int
	Type         int
	Flags        int
	IP           net.IP
	HardwareAddr net.HardwareAddr
}

// String returns $ip/$hwaddr $label
func (neigh *Neigh) String() string {
	return fmt.Sprintf("%d %s %s", neigh.LinkIndex, neigh.IP, neigh.HardwareAddr)
}
