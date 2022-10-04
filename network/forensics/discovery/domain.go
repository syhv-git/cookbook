package discovery

import (
	"github.com/syhv-git/cookbook/cmd"
	"net"
)

func MXFromDomain(v bool, domain string) (hosts []string, pref []uint16) {
	if domain == "" {
		cmd.Fatal("## No domain name provided")
	}
	cmd.Log(v, "- Looking up MX records for: %s", domain)
	mx, err := net.LookupMX(domain)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	for _, x := range mx {
		hosts, pref = append(hosts, x.Host), append(pref, x.Pref)
	}
	return
}

func NSFromDomain(v bool, domain string) (hosts []string) {
	if domain == "" {
		cmd.Fatal("## No domain name provided")
	}
	cmd.Log(v, "- Looking up name servers for: %s", domain)
	ns, err := net.LookupNS(domain)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	for _, x := range ns {
		hosts = append(hosts, x.Host)
	}
	return
}
