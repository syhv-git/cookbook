package discovery

import (
	"github.com/syhv-git/cookbook/cmd"
	"net"
)

func HostnamesFromIP(v bool, ip string) []string {
	if ip == "" {
		cmd.Fatal("## No IP address provided")
	}
	p := net.ParseIP(ip)

	cmd.Log(v, "- Looking up hostnames for: %s", p.String())
	host, err := net.LookupAddr(p.String())
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	return host
}

func IPsFromHostname(v bool, host string) (ipRes []string) {
	if host == "" {
		cmd.Fatal("## No hostname provided")
	}
	cmd.Log(v, "- Looking up IP addresses associated with: %s", host)

	ips, err := net.LookupIP(host)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	for _, ip := range ips {
		ipRes = append(ipRes, ip.String())
	}
	return
}
