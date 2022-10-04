package discovery

import (
	"github.com/google/gopacket/pcap"
	"github.com/syhv-git/cookbook/cmd"
)

func LogNetworkDevices() {
	v := true
	devs, err := pcap.FindAllDevs()
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	for _, x := range devs {
		cmd.Log(v, "- Network device name: %s", x.Name)
		cmd.Log(v, "- Network device description: %s", x.Description)
		cmd.Log(v, "- Network device flags: %v", x.Flags)
		for _, a := range x.Addresses {
			cmd.Log(v, "- IP address: %v", a.IP)
			cmd.Log(v, "- Subnet Mask: %v", a.Netmask)
		}
	}
}
