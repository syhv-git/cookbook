package forensics

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/syhv-git/cookbook/cmd"
	"time"
)

func CapturePackets(v, prom bool, device string, packetLen int32, timeout time.Duration) {
	if device == "" {
		cmd.Fatal("## No device name given")
	}

	handle, err := pcap.OpenLive(device, packetLen, prom, timeout)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer handle.Close()
	capture(v, 0, handle)
}

func CaptureNPackets(v, prom bool, device string, packetLen, n int32, timeout time.Duration) {
	if device == "" {
		cmd.Fatal("## No device name given")
	}

	handle, err := pcap.OpenLive(device, packetLen, prom, timeout)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer handle.Close()
	capture(v, n, handle)
}

func CapturePacketsFiltered(v, prom bool, device, filter string, packetLen int32, timeout time.Duration) {
	if device == "" {
		cmd.Fatal("## No device name given")
	}

	handle, err := pcap.OpenLive(device, packetLen, prom, timeout)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer handle.Close()

	if err = handle.SetBPFFilter(filter); err != nil {
		cmd.Fatal("## " + err.Error())
	}
	cmd.Log(v, "- Capturing packets based on the filter: %s", filter)
	capture(v, 0, handle)
}

func CaptureNPacketsFiltered(v, prom bool, device, filter string, packetLen, n int32, timeout time.Duration) {
	if device == "" {
		cmd.Fatal("## No device name given")
	}

	handle, err := pcap.OpenLive(device, packetLen, prom, timeout)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer handle.Close()

	if err = handle.SetBPFFilter(filter); err != nil {
		cmd.Fatal("## " + err.Error())
	}
	cmd.Log(v, "- Capturing packets based on the filter: %s", filter)
	capture(v, n, handle)
}

func capture(v bool, max int32, handle *pcap.Handle) {
	src := gopacket.NewPacketSource(handle, handle.LinkType())
	count := int32(0)
	for packet := range src.Packets() {
		cmd.Log(v, "- Packet captured: %v", packet)
		count++
		if count == max {
			break
		}
	}
}
