package forensics

import (
	"github.com/syhv-git/cookbook/cmd"
	"net"
)

func HostnamesFromIP(v bool, ip string) []string {
	p := net.ParseIP(ip)
	if p == nil {
		cmd.Fatal("## No IP address provided")
	}

	cmd.Log(v, "- Looking up hostnames for: %s", p.String())
	host, err := net.LookupAddr(p.String())
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	return host
}

func IPsFromHostname(v bool, host string) (res []string) {
	cmd.Log(v, "- Looking up IP addresses associated with: %s", host)

	ips, err := net.LookupIP(host)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	for _, ip := range ips {
		res = append(res, ip.String())
	}
	return
}
