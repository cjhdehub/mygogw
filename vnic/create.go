package vnic

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"net"
	"strings"
)

func isVrfExist(name string) (state bool, vrf *netlink.Vrf) {
	link, err := netlink.LinkByName(name)
	if err != nil {
		if strings.Contains(err.Error(), "Link not found") {
			return false, nil
		} else {
			//logger.Errorf("find vrf:%s err:%v", name, err)
			return false, nil
		}
	}
	var ok bool
	vrf, ok = link.(*netlink.Vrf)
	if ok {
		return true, vrf
	}

	return false, nil
}

func CreateVrf(vrfName string, tableId uint32) (err error) {
	exist, _ := isVrfExist(vrfName)
	if exist {
		return fmt.Errorf("vrf:%s already exist", vrfName)
	}

	vrf := &netlink.Vrf{
		LinkAttrs: netlink.LinkAttrs{
			Name:  vrfName,
			Flags: net.FlagUp,
		},
		Table: tableId,
	}

	err = netlink.LinkAdd(vrf)

	return
}

func isVethExist(vethName string) bool {
	_, err := netlink.LinkByName(vethName)
	if err != nil {
		return false
	}
	return true
}

func CreateVlan(vlanName, vlanIfIp, parentName string) error {
	if isVethExist(vlanName) {
		return fmt.Errorf("create failed vlan:%s already exist", vlanName)
	}

	parentLink, err := netlink.LinkByName(parentName)
	if err != nil {
		err = fmt.Errorf("find parentLink:%s err:%v", parentName, err)
		return err
	}

	vlan := &netlink.Vlan{
		LinkAttrs: netlink.LinkAttrs{
			Name:        vlanName,
			Flags:       net.FlagUp,
			ParentIndex: parentLink.Attrs().Index,
		},
		VlanId: 174,
	}

	err = netlink.LinkAdd(vlan)
	if err != nil {
		err = fmt.Errorf("LinkAdd err:%v", err)
		return err
	}

	var addr *netlink.Addr
	addr, err = netlink.ParseAddr(vlanIfIp)
	if err != nil {
		err = fmt.Errorf("ParseAddr err:%v", err)
		return err
	}

	err = netlink.AddrAdd(vlan, addr)
	if err != nil {
		err = fmt.Errorf("AddrAdd err:%v", err)
		return err
	}

	//err = netlink.LinkSetUp(vlan)
	//if err != nil {
	//	err = fmt.Errorf("LinkSetUp err:%v", err)
	//	return err
	//}

	return nil
}
