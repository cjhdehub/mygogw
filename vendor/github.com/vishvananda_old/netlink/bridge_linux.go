package netlink

import (
	"fmt"
	"syscall"

	"github.com/vishvananda/netlink/nl"
)

// BridgeVlanList gets a map of device id to bridge vlan infos.
// Equivalent to: `bridge vlan show`
func BridgeVlanList() (map[int32][]*nl.BridgeVlanInfo, error) {
	return pkgHandle.BridgeVlanList()
}

// BridgeVlanList gets a map of device id to bridge vlan infos.
// Equivalent to: `bridge vlan show`
func (h *Handle) BridgeVlanList() (map[int32][]*nl.BridgeVlanInfo, error) {
	req := h.newNetlinkRequest(syscall.RTM_GETLINK, syscall.NLM_F_DUMP)
	msg := nl.NewIfInfomsg(syscall.AF_BRIDGE)
	req.AddData(msg)
	req.AddData(nl.NewRtAttr(nl.IFLA_EXT_MASK, nl.Uint32Attr(uint32(nl.RTEXT_FILTER_BRVLAN))))

	msgs, err := req.Execute(syscall.NETLINK_ROUTE, syscall.RTM_NEWLINK)
	if err != nil {
		return nil, err
	}
	ret := make(map[int32][]*nl.BridgeVlanInfo)
	for _, m := range msgs {
		msg := nl.DeserializeIfInfomsg(m)

		attrs, err := nl.ParseRouteAttr(m[msg.Len():])
		if err != nil {
			return nil, err
		}
		for _, attr := range attrs {
			switch attr.Attr.Type {
			case nl.IFLA_AF_SPEC:
				//nested attr
				nestAttrs, err := nl.ParseRouteAttr(attr.Value)
				if err != nil {
					return nil, fmt.Errorf("failed to parse nested attr %v", err)
				}
				for _, nestAttr := range nestAttrs {
					switch nestAttr.Attr.Type {
					case nl.IFLA_BRIDGE_VLAN_INFO:
						vlanInfo := nl.DeserializeBridgeVlanInfo(nestAttr.Value)
						ret[msg.Index] = append(ret[msg.Index], vlanInfo)
					}
				}
			}
		}
	}
	return ret, nil
}
func BridgeVlanTunnelAdd(link Link, vid uint16, tunId int) error {
	return pkgHandle.bridgeVlanTunModify(syscall.RTM_SETLINK, link, vid, tunId)
}

func BridgeVlanTunnelDel(link Link, vid uint16, tunId int) error {
	return pkgHandle.bridgeVlanTunModify(syscall.RTM_DELLINK, link, vid, tunId)
}

// BridgeVlanAdd adds a new vlan filter entry
// Equivalent to: `bridge vlan add dev DEV vid VID [ pvid ] [ untagged ] [ self ] [ master ]`
func BridgeVlanAdd(link Link, vid uint16, pvid, untagged, self, master bool) error {
	return pkgHandle.BridgeVlanAdd(link, vid, pvid, untagged, self, master)
}

// BridgeVlanAdd adds a new vlan filter entry
// Equivalent to: `bridge vlan add dev DEV vid VID [ pvid ] [ untagged ] [ self ] [ master ]`
func (h *Handle) BridgeVlanAdd(link Link, vid uint16, pvid, untagged, self, master bool) error {
	return h.bridgeVlanModify(syscall.RTM_SETLINK, link, vid, pvid, untagged, self, master)
}

// BridgeVlanDel adds a new vlan filter entry
// Equivalent to: `bridge vlan del dev DEV vid VID [ pvid ] [ untagged ] [ self ] [ master ]`
func BridgeVlanDel(link Link, vid uint16, pvid, untagged, self, master bool) error {
	return pkgHandle.BridgeVlanDel(link, vid, pvid, untagged, self, master)
}

// BridgeVlanDel adds a new vlan filter entry
// Equivalent to: `bridge vlan del dev DEV vid VID [ pvid ] [ untagged ] [ self ] [ master ]`
func (h *Handle) BridgeVlanDel(link Link, vid uint16, pvid, untagged, self, master bool) error {
	return h.bridgeVlanModify(syscall.RTM_DELLINK, link, vid, pvid, untagged, self, master)
}

func BridgeVlanVniList(link Link) (map[int]int, error) {
	return pkgHandle.BridgeVlanVniList(link)
}

func (h *Handle) BridgeVlanVniList(link Link) (map[int]int, error) {
	req := h.newNetlinkRequest(syscall.RTM_GETLINK, syscall.NLM_F_DUMP)
	msg := nl.NewIfInfomsg(syscall.AF_BRIDGE)
	base := link.Attrs()
	h.ensureIndex(base)
	msg.Index = int32(base.Index)

	req.AddData(msg)
	req.AddData(nl.NewRtAttr(nl.IFLA_EXT_MASK, nl.Uint32Attr(uint32(nl.RTEXT_FILTER_BRVLAN))))

	msgs, err := req.Execute(syscall.NETLINK_ROUTE, syscall.RTM_NEWLINK)
	if err != nil {
		return nil, err
	}
	ret := make(map[int]int)
	for _, m := range msgs {
		msg := nl.DeserializeIfInfomsg(m)

		attrs, err := nl.ParseRouteAttr(m[msg.Len():])
		if err != nil {
			return nil, err
		}
		for _, attr := range attrs {
			switch attr.Attr.Type {
			case nl.IFLA_AF_SPEC:
				//nested attr
				nestAttrs, err := nl.ParseRouteAttr(attr.Value)
				if err != nil {
					return nil, fmt.Errorf("failed to parse nested attr %v", err)
				}
				for _, nestAttr := range nestAttrs {
					switch nestAttr.Attr.Type {
					case nl.IFLA_BRIDGE_VLAN_TUNNEL_INFO:
						vlanVni, err2 := nl.ParseRouteAttr(nestAttr.Value)
						if err != nil {
							return nil, fmt.Errorf("failed to parse IFLA_BRIDGE_VLAN_TUNNEL_INFO nested attr %v", err2)
						}
						var vlan int
						var vni int
						for _, val := range vlanVni {
							switch val.Attr.Type {
							case 1:
								vni = int(native.Uint32(val.Value[0:4]))
							case 2:
								vlan = int(native.Uint16(val.Value[0:2]))

							}
						}
						ret[vlan] = vni

					}
				}
			}
		}
	}
	return ret, nil
}

func (h *Handle) bridgeVlanTunModify(cmd int, link Link, vid uint16, tunId int) error {
	base := link.Attrs()
	h.ensureIndex(base)
	req := h.newNetlinkRequest(cmd, syscall.NLM_F_ACK)

	msg := nl.NewIfInfomsg(syscall.AF_BRIDGE)
	msg.Index = int32(base.Index)
	req.AddData(msg)

	br := nl.NewRtAttr(nl.IFLA_AF_SPEC, nil)

	rt := nl.NewRtAttrChild(br, nl.IFLA_BRIDGE_VLAN_TUNNEL_INFO, nil)

	nl.NewRtAttrChild(rt, 1, nl.Uint32Attr(uint32(tunId)))
	nl.NewRtAttrChild(rt, 2, nl.Uint16Attr(vid))

	req.AddData(br)
	_, err := req.Execute(syscall.NETLINK_ROUTE, 0)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handle) bridgeVlanModify(cmd int, link Link, vid uint16, pvid, untagged, self, master bool) error {
	base := link.Attrs()
	h.ensureIndex(base)
	req := h.newNetlinkRequest(cmd, syscall.NLM_F_ACK)

	msg := nl.NewIfInfomsg(syscall.AF_BRIDGE)
	msg.Index = int32(base.Index)
	req.AddData(msg)

	br := nl.NewRtAttr(nl.IFLA_AF_SPEC, nil)
	var flags uint16
	if self {
		flags |= nl.BRIDGE_FLAGS_SELF
	}
	if master {
		flags |= nl.BRIDGE_FLAGS_MASTER
	}
	if flags > 0 {
		nl.NewRtAttrChild(br, nl.IFLA_BRIDGE_FLAGS, nl.Uint16Attr(flags))
	}
	vlanInfo := &nl.BridgeVlanInfo{Vid: vid}
	if pvid {
		vlanInfo.Flags |= nl.BRIDGE_VLAN_INFO_PVID
	}
	if untagged {
		vlanInfo.Flags |= nl.BRIDGE_VLAN_INFO_UNTAGGED
	}
	nl.NewRtAttrChild(br, nl.IFLA_BRIDGE_VLAN_INFO, vlanInfo.Serialize())
	req.AddData(br)
	_, err := req.Execute(syscall.NETLINK_ROUTE, 0)
	if err != nil {
		return err
	}
	return nil
}
